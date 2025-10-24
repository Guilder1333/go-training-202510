package logic

import (
	"errors"
	"fmt"
	"math/rand"
)

type userServiceStubImpl struct {
}

func NewUserServiceStub() UserService {
	return &userServiceStubImpl{}
}

func (u *userServiceStubImpl) GetUserById(id int) (*User, error) {
	if id == 404 {
		return nil, fmt.Errorf("special case handling for 404 user id: %w", ErrUserNotFound)
	}
	if id == 500 {
		return nil, errors.New("failed to get user")
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
		return 0, errors.New("failed to create user")
	}

	return 1000 + rand.Intn(1000), nil
}

func (u *userServiceStubImpl) DeleteUser(id int) error {
	if id == 404 {
		return fmt.Errorf("special case handling for 404 user id: %w", ErrUserNotFound)
	}
	if id == 500 {
		return errors.New("failed to delete user")
	}
	return nil
}
