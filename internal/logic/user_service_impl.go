package logic

import (
	"errors"
	"fmt"
	"handsongo/internal/dal"
	"handsongo/internal/statuserror"
)

type UserServiceImpl struct {
	repository dal.UserRepository
}

func NewUserService(repository dal.UserRepository) UserService {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (s *UserServiceImpl) GetUserById(id int) (*User, error) {
	return nil, errors.New("not implemented")
}
func (s *UserServiceImpl) CreateUser(user *User) (int, error) {
	return 0, errors.New("not implemented")
}

func (s *UserServiceImpl) DeleteUser(id int) error {
	exists, err := s.repository.CheckUserById(id)
	if err != nil {
		return statuserror.SetErrorMessage("failed to delete user",
			fmt.Errorf("failed to check user existance in repository: %w", err))
	}

	if !exists {
		return statuserror.SetErrorMessage("user not found",
			statuserror.SetStatusError(statuserror.ErrorKindNotFound,
				errors.New("user not found in repository")))
	}

	err = s.repository.DeleteUser(id)
	if err != nil {
		return statuserror.SetErrorMessage("failed to delete user",
			fmt.Errorf("failed to delete user from repository: %w", err))
	}
	return nil
}
