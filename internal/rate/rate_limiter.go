package rate

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var limiter = redis_rate.NewLimiter(client)

// func RateLimit(next http.Handler) http.Handler {
func RateLimit(next http.Handler, ipMaxRequest int32, tokenMaxRequest int32, apiKey string, redisHost string, redisPort int32, redisPassword string, redisDb string, webServerPort string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if r.Header.Get("API_KEY") != "" {
		if r.Header.Get("API_KEY") == apiKey {
			rateLimitToken(next).ServeHTTP(w, r)
		} else {
			rateLimitIp(next).ServeHTTP(w, r)
		}
	})
}

func rateLimitIp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		// fmt.Print(redis_rate.PerSecond(10), "\n")
		// fmt.Print("Por ip \n")
		//res, err := limiter.Allow(r.Context(), ip, redis_rate.PerSecond(10))
		res, err := limiter.Allow(r.Context(), ip, redis_rate.PerSecond(10))
		if err != nil {
			tooManyRequestErrors(w)
			return
		}

		if res.Allowed > 0 {
			next.ServeHTTP(w, r)
		} else {
			tooManyRequestErrors(w)
		}
	})
}

func rateLimitToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("API_KEY")
		// fmt.Print(redis_rate.PerSecond(100), "\n")
		// fmt.Print("Por Token \n")
		// fmt.Print(token, "\n")
		res, err := limiter.Allow(r.Context(), token, redis_rate.PerSecond(100))
		if err != nil {
			tooManyRequestErrors(w)
			return
		}

		if res.Allowed > 0 {
			next.ServeHTTP(w, r)
		} else {
			tooManyRequestErrors(w)
		}
	})
}

func tooManyRequestErrors(w http.ResponseWriter) {
	http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
}
