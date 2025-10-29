package presentation

import (
	"handsongo/internal/statuserror"
	"net/http"

	"github.com/rs/zerolog/log"
)

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error

func wrapErrorResponse(handlerFn HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handlerFn(w, r)
		if err != nil {
			log.Error().Err(err).Msg("Handler execution has failed")

			msg := "error during handling response"
			statusCode := 500
			kind := statuserror.GetErrorKind(err)
			switch kind {
			case statuserror.ErrorKindInvalidRequest:
				msg = "request parameters validation failed"
				statusCode = 400
			case statuserror.ErrorKindNotFound:
				msg = "user not found"
				statusCode = 404
			}

			errormsg, ok := statuserror.GetErrorMessage(err)
			if ok {
				msg = errormsg
			}

			w.WriteHeader(statusCode)
			w.Write([]byte(msg))
		}
	}
}
