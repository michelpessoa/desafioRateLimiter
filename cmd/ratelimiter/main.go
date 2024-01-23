package main

import (
	"fmt"
	"net/http"

	"github.com/michelpessoa/desafioRateLimiter/configs"
	"github.com/michelpessoa/desafioRateLimiter/internal/rate"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	// Use o rate limiter como middlewares
	limitedMux := rate.RateLimit(mux, configs.IpMaxRequest, configs.TokenMaxRequest, configs.ApiKey, configs.RedisHost, configs.RedisPort, configs.RedisPassword, configs.RedisDb, configs.WebServerPort)

	http.ListenAndServe(":8080", limitedMux)
}
