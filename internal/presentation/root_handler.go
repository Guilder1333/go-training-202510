package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRootHandler(userContoller *UserController) http.Handler {
	router := chi.NewRouter()
	router.Get("/user", wrapErrorResponse(userContoller.GetUserById))
	router.Post("/user", wrapErrorResponse(userContoller.CreateUser))
	router.Delete("/user", wrapErrorResponse(userContoller.DeleteUserById))

	return router
}
