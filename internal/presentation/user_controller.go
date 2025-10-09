package presentation

import (
	"encoding/json"
	"errors"
	"handsongo/internal/logic"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

type getUserByIDResponse struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Age             int    `json:"age"`
	PhoneNumber     string `json:"phoneNumber"`
	IsPhoneVerified bool   `json:"isPhoneVerified"`
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

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	// 1. Get ID from request (query string)
	//	1.1. if this is invalid -> return 400 response
	getUserRequest, err := u.validator.ValidateGetUserById(r)
	if err != nil {
		log.Warn().Err(err).Msg("Passed user get by id parameters were invalid")
		w.WriteHeader(400)
		w.Write([]byte("request parameters validation failed"))
		return
	}

	// 2. Pass ID to business logic layer & get back a user
	//    2.1. If user is not found -> return 404 response
	user, err := u.userService.GetUserById(getUserRequest.Id)
	if err != nil {
		if errors.Is(err, logic.ErrUserNotFound) {
			log.Warn().Err(err).Msg("Requesting non-existant user " + strconv.Itoa(getUserRequest.Id))
			w.WriteHeader(404)
			w.Write([]byte("user not found"))
			return
		}

		log.Error().Err(err).Msg("Failed to get user")
		w.WriteHeader(500)
		w.Write([]byte("failed to get user"))
		return
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
		log.Error().Err(err).Msg("User response serialization failed")
		w.WriteHeader(500)
		w.Write([]byte("failed to make response"))
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
