package presentation

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var errIdParameterIsMissing = errors.New("missing ID parameter")

type UserValidatorImpl struct {
}

func (v *UserValidatorImpl) ValidateGetUserById(r *http.Request) (UserGetResponse, error) {
	str := r.URL.Query().Get("id")
	if str == "" {
		return UserGetResponse{}, errIdParameterIsMissing
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return UserGetResponse{}, fmt.Errorf("failed to convert id to integer: %w", err)
	}

	return UserGetResponse{Id: id}, nil
}
