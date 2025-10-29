package logic

import (
	"errors"
	"handsongo/internal/statuserror"
	"math/rand"
)

type userServiceStubImpl struct {
}

func NewUserServiceStub() UserService {
	return &userServiceStubImpl{}
}

func (u *userServiceStubImpl) GetUserById(id int) (*User, error) {
	if id == 404 {
		return nil,
			statuserror.SetErrorMessage("user not found",
				statuserror.SetStatusError(
					statuserror.ErrorKindNotFound,
					errors.New("special case handling for 404 user id")))
	}
	if id == 500 {
		return nil, statuserror.SetErrorMessage("failed to get user", errors.New("db error"))
	}

	return &User{
		Id:              id,
		FirstName:       "Rakuten",
		LastName:        "Taro",
		Age:             24,
		PhoneNumber:     "12346567",
		IsPhoneVerified: true,
	}, nil
}

func (u *userServiceStubImpl) CreateUser(user *User) (int, error) {
	if user.FirstName == "Invalid" {
		return 0, statuserror.SetErrorMessage("failed to create user", errors.New("db error"))
	}

	return 1000 + rand.Intn(1000), nil
}

func (u *userServiceStubImpl) DeleteUser(id int) error {
	if id == 404 {
		statuserror.SetErrorMessage("user not found",
			statuserror.SetStatusError(
				statuserror.ErrorKindNotFound,
				errors.New("special case handling for 404 user id")))
	}
	if id == 500 {
		return statuserror.SetErrorMessage("failed to delete user", errors.New("db error"))
	}
	return nil
}
