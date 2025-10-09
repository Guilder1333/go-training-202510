package logic

import "errors"

var ErrUserNotFound = errors.New("user not found")

type GetUserById struct {
	Id              int
	FirstName       string
	LastName        string
	Age             int
	PhoneNumber     string
	IsPhoneVerified bool
}

type UserService interface {
	GetUserById(id int) (*GetUserById, error)
}
