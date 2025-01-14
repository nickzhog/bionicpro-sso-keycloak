package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Get("/reports", handleReports)

		r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
			// Cors or Unexpected
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(200)
		})
	})
	fmt.Println("start")
	fmt.Println(http.ListenAndServe(":8000", r))
	fmt.Println("stop")
}

func handleReports(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	fmt.Println("token:" + token)
	token = strings.TrimPrefix(token, "Bearer ")

	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if token == "" {
		http.Error(w, "need to use token", http.StatusUnauthorized)
		return
	}

	isAllowed, err := auth(token)
	if err != nil {
		http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
		return
	}
	if !*isAllowed {
		http.Error(w, "you role cant access here", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("report123"))
}
