package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/michelpessoa/desafioRateLimiter/internal/storage"
	"github.com/redis/go-redis/v9"
)

var _ storage.Storage = (*RedisStorage)(nil)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(user string, password string, host string, port string, dbName string) *RedisStorage {
	database, err := strconv.Atoi(dbName)

	if err != nil {
		database = 0
	}

	return &RedisStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       database,
		}),
	}
}

func (s *RedisStorage) Increment(ctx context.Context, key string, ttl int) (int, error) {
	pipe := s.client.Pipeline()

	pipe.Exists(ctx, key)
	pipe.Incr(ctx, key)

	counter, err := pipe.Exec(ctx) // Execute the pipeline

	if err != nil {
		return 0, err
	}

	if len(counter) > 0 && counter[0].(*redis.IntCmd).Val() == 0 { // Key didn't exist
		err = pipe.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
	}

	if err != nil {
		return 0, err
	}

	_, err = pipe.Exec(ctx) // Ensure pipeline execution
	if err != nil {
		return 0, err
	}

	return int(counter[1].(*redis.IntCmd).Val()), nil
}

func (s *RedisStorage) Get(ctx context.Context, key string) (interface{}, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *RedisStorage) Set(ctx context.Context, key string, ttl int) error {
	var err error
	pipe := s.client.Pipeline()
	pipe.Set(ctx, key, true, 0)

	if ttl != 0 {
		err = pipe.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) Exists(ctx context.Context, key string) (bool, error) {
	result, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return result == 1, nil
}

func (s *RedisStorage) IsBlocked(ctx context.Context, key string) (bool, error) {
	return s.Exists(ctx, key)
}
