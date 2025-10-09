package logic

import (
	"errors"
	"fmt"
)

type userServiceStubImpl struct {
}

func NewUserServiceStub() UserService {
	return &userServiceStubImpl{}
}

func (u *userServiceStubImpl) GetUserById(id int) (*GetUserById, error) {
	if id == 404 {
		return nil, fmt.Errorf("special case handling for 404 user id: %w", ErrUserNotFound)
	}
	if id == 500 {
		return nil, errors.New("failed to get user")
	}

	return &GetUserById{
		Id:              id,
		FirstName:       "Rakuten",
		LastName:        "Taro",
		Age:             24,
		PhoneNumber:     "12346567",
		IsPhoneVerified: true,
	}, nil
}
