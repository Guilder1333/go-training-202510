package statuserror

import "errors"

type ErrorMessage struct {
	msg     string
	wrapped error
}

func SetErrorMessage(msg string, err error) error {
	return &ErrorMessage{
		msg:     msg,
		wrapped: err,
	}
}

func GetErrorMessage(err error) (msg string, ok bool) {
	var errormessage *ErrorMessage
	if errors.As(err, &errormessage) {
		return errormessage.msg, true
	}
	return "", false
}

func (e *ErrorMessage) Error() string {
	if e.wrapped != nil {
		return ""
	}
	return e.wrapped.Error()
}

func (e *ErrorMessage) Unwrap() error {
	return e.wrapped
}
