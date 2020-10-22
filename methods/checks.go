package methods

import (
	auth "github.com/IT108/achieve-auth-go/auth"
	models "github.com/IT108/achieve-models-go"
	"net/http"
)

func IsRegistered(request models.IsRegisteredRequest) models.IsRegisteredResponse {
	response := models.IsRegisteredResponse{
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

func IsEmailRegistered(request models.IsEmailRegisteredRequest) models.IsEmailResponse {
	result := models.IsEmailResponse{
		Request:           request.Request,
		ResponseCode:      http.StatusOK,
		IsEmailRegistered: false,
		Error:             "",
	}

	ok, err := auth.IsEmailAvailable(request.Email)
	if !ok {
		result.ResponseCode = http.StatusConflict
		result.Error = err
		result.IsEmailRegistered = true
	}

	return result
}

func IsUserRegistered(request models.IsUserRegisteredRequest) models.IsUserRegisteredResponse {
	result := models.IsUserRegisteredResponse{
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
