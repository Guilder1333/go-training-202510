package presentation

import (
	"errors"
	"net/http"
)

var ErrInvalidRequest = errors.New("request parameters are invalid")

type UserGetRequestBody struct {
	Id int
}

type CreateUserRequestBody struct {
	FirstName       string
	LastName        string
	Age             int
	PhoneNumber     string
	IsPhoneVerified bool
}

type UserDeleteRequestBody struct {
	Id int
}

type UserValidator interface {
	ValidateGetUserById(r *http.Request) (UserGetRequestBody, error)
	ValidateCreateUser(r *http.Request) (*CreateUserRequestBody, error)
	ValidateDeleteUser(r *http.Request) (UserDeleteRequestBody, error)
}
