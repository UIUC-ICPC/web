package main

import (
	"log"
	"net/http"
	"sigicpc-backend/internal/handlers"
	"sigicpc-backend/internal/middleware"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/members", handlers.GetMembers)
	mux.HandleFunc("GET /api/posts", handlers.GetPosts)

	log.Println("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(middleware.RateLimit(mux))))
}
