package presentation

import "net/http"

type UserGetResponse struct {
	Id int
}

type UserValidator interface {
	ValidateGetUserById(r *http.Request) (UserGetResponse, error)
}
