package presentation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var errIdParameterIsMissing = errors.New("missing ID parameter")

type UserValidatorImpl struct {
}

func NewUserValidator() UserValidator {
	return &UserValidatorImpl{}
}

func (v *UserValidatorImpl) ValidateGetUserById(r *http.Request) (UserGetRequestBody, error) {
	str := r.URL.Query().Get("id")
	if str == "" {
		return UserGetRequestBody{}, errIdParameterIsMissing
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return UserGetRequestBody{}, fmt.Errorf("failed to convert id to integer: %w", err)
	}

	return UserGetRequestBody{Id: id}, nil
}

type createUserRequestJson struct {
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Age             *int    `json:"age"`
	PhoneNumber     *string `json:"phoneNumber"`
	IsPhoneVerified *bool   `json:"isPhoneVerified"`
}

func (v *UserValidatorImpl) ValidateCreateUser(r *http.Request) (*CreateUserRequestBody, error) {
	// Parse json
	var body createUserRequestJson
	err := parseBody(nil, r.Body, &body)
	// Check required fields.
	err = required(err, body.FirstName, "FirstName")
	err = required(err, body.LastName, "LastName")
	err = required(err, body.Age, "Age")
	err = required(err, body.PhoneNumber, "PhoneNumber")
	err = required(err, body.IsPhoneVerified, "IsPhoneVerified")
	// Check fields constrains.

	err = stringLength(err, body.FirstName, 1, 100, "FirstName")
	err = stringLength(err, body.LastName, 1, 100, "LastName")
	err = stringLength(err, body.PhoneNumber, 10, 25, "PhoneNumber")
	err = intValue(err, body.Age, 0, 200, "Age")
	if err != nil {
		return nil, err
	}

	// output conversion
	request := CreateUserRequestBody{
		FirstName:       *body.FirstName,
		LastName:        *body.LastName,
		Age:             *body.Age,
		PhoneNumber:     *body.PhoneNumber,
		IsPhoneVerified: *body.IsPhoneVerified,
	}

	return &request, nil
}

func parseBody(err error, body io.ReadCloser, value any) error {
	if err != nil {
		return err
	}
	err = json.NewDecoder(body).Decode(&value)
	if err != nil {
		return fmt.Errorf("%w: failed to parse request body: %w",
			ErrInvalidRequest, err)
	}
	return nil
}

func required[T any](err error, value *T, valueName string) error {
	if err != nil {
		return err
	}
	if value == nil {
		return fmt.Errorf("%w: missing field '%s' in request body",
			ErrInvalidRequest, valueName)
	}
	return nil
}

func stringLength(err error, str *string, min int, max int, valueName string) error {
	if err != nil || str == nil {
		return err
	}

	strLen := len(*str)
	if strLen < min || strLen > max {
		return fmt.Errorf("%w: string '%s' length does not fit into constraints [%d, %d]", ErrInvalidRequest, valueName, min, max)
	}
	return nil
}

func intValue(err error, value *int, min int, max int, valueName string) error {
	if err != nil || value == nil {
		return err
	}

	v := *value
	if v < min || v > max {
		return fmt.Errorf("%w: value '%s' does not fit into constraints [%d, %d]", ErrInvalidRequest, valueName, min, max)
	}
	return nil
}

func (v *UserValidatorImpl) ValidateDeleteUser(r *http.Request) (UserDeleteRequestBody, error) {
	str := r.URL.Query().Get("id")
	if str == "" {
		return UserDeleteRequestBody{}, errIdParameterIsMissing
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return UserDeleteRequestBody{}, fmt.Errorf("failed to convert id to integer: %w", err)
	}

	return UserDeleteRequestBody{Id: id}, nil
}
