package methods

import (
	"context"
	"errors"
	auth "github.com/IT108/achieve-auth-go/auth"
	auth2 "github.com/IT108/achieve-models-go/auth"
	"log"
	"net/http"
)

func IsRegistered(request auth2.IsRegisteredRequest) auth2.IsRegisteredResponse {
	response := auth2.IsRegisteredResponse{
		Request:              request.Request,
		ResponseCode:         http.StatusOK,
		IsEmailRegistered:    false,
		IsUsernameRegistered: false,
		Error:                "",
	}

	email, err := auth.IsEmailAvailable(request.Email)
	if !email {
		response.IsEmailRegistered = true
		response.ResponseCode = http.StatusConflict
		response.Error = err
	}

	username, err := auth.IsUsernameAvailable(request.Email)
	if !username {
		response.IsUsernameRegistered = true
		response.ResponseCode = http.StatusConflict
		response.Error = err
	}

	if !username && !email {
		response.Error = "Email and username are registered"
	}

	return response
}

func (s *Server) IsEmailRegistered(ctx context.Context, request *auth.IsEmailRequest) (*auth.IsEmailResponse, error) {
	log.Print("Is email")

	result := auth.IsEmailResponse{
		IsEmailRegistered: false,
	}

	ok, _ := auth.IsEmailAvailable(request.Email)
	if !ok {
		result.IsEmailRegistered = true
		return &result, errors.New("email already registered")
	}

	return &result, nil
}

func IsUserRegistered(request auth2.IsUserRegisteredRequest) auth2.IsUserRegisteredResponse {
	result := auth2.IsUserRegisteredResponse{
		Request:              request.Request,
		ResponseCode:         http.StatusOK,
		IsUsernameRegistered: false,
		Error:                "",
	}

	ok, err := auth.IsUsernameAvailable(request.Username)
	if !ok {
		result.ResponseCode = http.StatusConflict
		result.Error = err
		result.IsUsernameRegistered = true
	}

	return result
}
