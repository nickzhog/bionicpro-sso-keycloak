package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		})
	})
	fmt.Println("start")
	fmt.Println(http.ListenAndServe(":8000", r))
	fmt.Println("stop")
}
