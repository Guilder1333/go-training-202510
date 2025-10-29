package userapp

import (
	"handsongo/internal/logic"
	"handsongo/internal/presentation"
	"net/http"
)

type userApplication struct {
	rootHandler http.Handler
}

func newUserApplication() (*userApplication, error) {
	var app userApplication

	validator := presentation.NewUserValidator()
	userService := logic.NewUserServiceStub()

	userController := presentation.NewUserController(
		validator,
		userService)

	app.rootHandler = presentation.NewRootHandler(userController)

	return &app, nil
}

func closeApplication(app *userApplication) {

}
