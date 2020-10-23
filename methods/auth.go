package methods

import (
	"github.com/IT108/achieve-auth-go/auth"
	auth_models "github.com/IT108/achieve-models-go/auth"
	"log"
	"net/http"
)

func Register(request auth_models.RegisterRequest) auth_models.RegisterResponse {

	ok, err := auth.Register(request.Username, request.Email, request.Password)

	response := auth_models.RegisterResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
	}

	if !ok {
		response.ResponseCode = http.StatusConflict
		response.Error = err
	}

	return response
}

func Authenticate(request auth_models.AuthenticateRequest) auth_models.AuthenticateResponse {
	ok, err := auth.Authenticate(request.Username, request.Password)

	response := auth_models.AuthenticateResponse{
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

func Authorize(request auth_models.AuthorizeRequest) auth_models.AuthorizeResponse {
	ok, groups, err := auth.Authorize(request.User)

	response := auth_models.AuthorizeResponse{
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

func SignIn(request auth_models.SigninRequest) auth_models.SigninResponse {
	ok, err := auth.Authenticate(request.Username, request.Password)

	response := auth_models.SigninResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
		Error:        "",
		Token:        auth_models.AuthToken{},
	}

	if !ok {
		response.ResponseCode = http.StatusUnauthorized
		response.Error = err
		return response
	}
	username := request.Username
	response.User = username

	ok, groups, err := auth.Authorize(username)

	if !ok {
		response.ResponseCode = http.StatusConflict
		response.Error = err
		return response
	}

	log.Println(groups)
	//TODO: custom claims

	ok, token, err := auth.GenerateToken(username)
	if !ok {
		response.Error = err
		response.ResponseCode = http.StatusInternalServerError
		return response
	}

	response.Token = token
	return response
}
