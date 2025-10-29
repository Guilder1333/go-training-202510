package statuserror

import "errors"

type Unwrappable interface {
	Unwrap() []error
}

type StatusError struct {
	kind    ErrorKind
	wrapped error
}

func SetStatusError(kind ErrorKind, err error) error {
	return &StatusError{
		kind:    kind,
		wrapped: err,
	}
}

func GetErrorKind(err error) ErrorKind {
	var statuserror *StatusError
	if errors.As(err, &statuserror) {
		return statuserror.kind
	}
	return ErrorKindUnknown
}

func (e *StatusError) Unwrap() error {
	return e.wrapped
}

func (e *StatusError) Error() string {
	if e.wrapped == nil {
		return ""
	}
	return e.wrapped.Error()
}
