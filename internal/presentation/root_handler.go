package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRootHandler(userContoller *UserController) http.Handler {
	router := chi.NewRouter()
	router.Get("/user", userContoller.GetUserById)
	router.Post("/user", userContoller.CreateUser)
	router.Delete("/user", userContoller.DeleteUserById)

	return router
}
