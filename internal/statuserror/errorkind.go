package statuserror

type ErrorKind int

const (
	ErrorKindUnknown ErrorKind = iota
	ErrorKindNotFound
	ErrorKindInvalidRequest
)
