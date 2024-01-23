package main

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mtx sync.Mutex

// Cria um novo rate limiter e adiciona o token do visitante ao mapa de visitantes.
func addVisitor(token string) *rate.Limiter {
	limiter := rate.NewLimiter(100, 1)
	mtx.Lock()
	visitors[token] = limiter
	mtx.Unlock()
	return limiter
}

// Obtém o rate limiter para o token do visitante atual. Se o visitante não existir no mapa de visitantes, chama a função addVisitor.
func getVisitor(token string) *rate.Limiter {
	mtx.Lock()
	limiter, exists := visitors[token]
	if !exists {
		mtx.Unlock()
		return addVisitor(token)
	}
	mtx.Unlock()
	return limiter
}

// Middleware que limita a taxa de solicitações
func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token") // substitua por sua lógica de extração de token
		limiter := getVisitor(token)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem-vindo!"))
	})

	// Limpa o mapa de visitantes a cada 5 minutos para evitar o crescimento do mapa indefinidamente.
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			mtx.Lock()
			visitors = make(map[string]*rate.Limiter)
			mtx.Unlock()
		}
	}()

	http.ListenAndServe(":8080", limit(mux))
}
