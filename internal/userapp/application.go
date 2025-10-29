package userapp

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func Run() {
	// Setup Application
	app, err := newUserApplication()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize application")
	}

	log.Info().Msg("Starting server")

	// Run our application on HTTP server
	runServer(app)

	log.Info().Msg("Terminating")

	// Shutdown HTTP server
	// Teardown application
	closeApplication(app)
}

func runServer(app *userApplication) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	errchan := make(chan error, 1)

	server := http.Server{
		Addr:    ":8080",
		Handler: app.rootHandler,
	}
	go listenAndServe(&server, errchan)
	log.Info().Msg("Listening to the server")

	select {
	case sig := <-sigchan:
		log.Info().Msg("Got signal" + sig.String())

		err := server.Shutdown(context.TODO())
		if err != nil {
			log.Error().Err(err).Msg("Correct shutdown of the HTTP server has failed")
		}
	case err := <-errchan:
		log.Error().Err(err).Msg("Failed to start HTTP server")
	}
}

func listenAndServe(server *http.Server, errchan chan error) {
	err := server.ListenAndServe()
	if err != nil {
		errchan <- err
	}
}
