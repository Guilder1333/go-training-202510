package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRootHandler() http.Handler {
	router := chi.NewRouter()
	router.Get("/user", nil)
	router.Post("/user", nil)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
		w.WriteHeader(200)
	})
}
