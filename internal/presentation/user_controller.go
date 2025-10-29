package presentation

import (
	"encoding/json"
	"fmt"
	"handsongo/internal/logic"
	"net/http"
)

type getUserByIDResponse struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Age             int    `json:"age"`
	PhoneNumber     string `json:"phoneNumber"`
	IsPhoneVerified bool   `json:"isPhoneVerified"`
}

type createUserResponse struct {
	Id int `json:"id"`
}

type UserController struct {
	validator   UserValidator
	userService logic.UserService
}

func NewUserController(validator UserValidator, userService logic.UserService) *UserController {
	return &UserController{
		validator:   validator,
		userService: userService,
	}
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) error {
	// 1. Get ID from request (query string)
	//	1.1. if this is invalid -> return 400 response
	getUserRequest, err := u.validator.ValidateGetUserById(r)
	if err != nil {
		return fmt.Errorf("get user by id request parameters validation failed: %w", err)
	}

	// 2. Pass ID to business logic layer & get back a user
	//    2.1. If user is not found -> return 404 response
	user, err := u.userService.GetUserById(getUserRequest.Id)
	if err != nil {
		return fmt.Errorf("failed to get user by id=%d: %w", getUserRequest.Id, err)
	}

	// 3. Map user model to user JSON response model
	userResponse := getUserByIDResponse{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Age:             user.Age,
		PhoneNumber:     user.PhoneNumber,
		IsPhoneVerified: user.IsPhoneVerified,
	}

	// 4. Serialize JSON response model and write 200 response with JSON
	//	body.
	response, err := json.Marshal(userResponse)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Write(response)
	return nil
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) error {
	// 1. Parse and validate request body
	body, err := u.validator.ValidateCreateUser(r)
	if err != nil {
		return fmt.Errorf("create user request parameters validation failed: %w", err)
	}
	// 2. Pass data to userservice
	user := logic.User{
		Id:              0,
		FirstName:       body.FirstName,
		LastName:        body.LastName,
		Age:             body.Age,
		PhoneNumber:     body.PhoneNumber,
		IsPhoneVerified: body.IsPhoneVerified,
	}
	id, err := u.userService.CreateUser(&user)
	if err != nil {
		return fmt.Errorf("failed to create user")
	}
	// 3. Response with the new user ID.
	response := createUserResponse{
		Id: id,
	}
	responseBody, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.WriteHeader(201)
	w.Write(responseBody)
	return nil
}

func (u *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) error {
	deleteUserRequest, err := u.validator.ValidateDeleteUser(r)
	if err != nil {
		return fmt.Errorf("delete user by id request parameters validation failed: %w", err)
	}

	err = u.userService.DeleteUser(deleteUserRequest.Id)
	if err != nil {
		return fmt.Errorf("failed to delete user by id %d: %w", deleteUserRequest.Id, err)
	}

	w.WriteHeader(204)
	return nil
}
