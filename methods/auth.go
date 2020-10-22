package methods

import (
	auth "github.com/IT108/achieve-auth-go/auth"
	models "github.com/IT108/achieve-models-go"
	"log"
	"net/http"
)

func Register(request models.RegisterRequest) models.RegisterResponse {

	ok, err := auth.Register(request.Username, request.Email, request.Password)

	response := models.RegisterResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
	}

	if !ok {
		response.ResponseCode = http.StatusConflict
		response.Error = err
	}

	return response
}

func Authenticate(request models.AuthenticateRequest) models.AuthenticateResponse {
	ok, err := auth.Authenticate(request.Username, request.Password)

	response := models.AuthenticateResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
	}

	if !ok {
		response.ResponseCode = http.StatusForbidden
		response.Error = err
	}

	log.Println(response.ResponseCode)
	return response
}

func Authorize(request models.AuthorizeRequest) models.AuthorizeResponse {
	ok, groups, err := auth.Authorize(request.User)

	response := models.AuthorizeResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
		Roles:        nil,
		Error:        "",
	}

	if !ok {
		response.ResponseCode = http.StatusConflict
		response.Error = err
		return response
	}
	response.Roles = groups
	return response
}
