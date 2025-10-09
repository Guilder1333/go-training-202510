package userapp

import (
	"handsongo/internal/presentation"
	"net/http"
)

type userApplication struct {
	rootHandler http.Handler
}

func newUserApplication() (*userApplication, error) {
	var app userApplication

	app.rootHandler = presentation.NewRootHandler()

	return &app, nil
}

func closeApplication(app *userApplication) {
	// TODO
}

// TODO
// Create those two files and their content
// Take a rest UNTIL 14.30
