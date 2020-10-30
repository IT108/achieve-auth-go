package methods

import (
	"context"
	"errors"
	"github.com/IT108/achieve-auth-go/auth"
	auth_models "github.com/IT108/achieve-models-go/auth"
	"github.com/golang/protobuf/ptypes"
	"log"
	"net/http"
)

type Server struct {
}

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


func (s *Server) SignIn(ctx context.Context, request *auth.SignInRequest) (*auth.Token, error) {
	ok, err := auth.Authenticate(request.Username, request.Password)

	response := auth.Token{
		Token:           "",
		TokenExpiration: nil,
	}

	if !ok {
		return &response, errors.New(err)
	}

	username := request.Username

	ok, groups, err := auth.Authorize(username)

	if !ok {

		return &response, errors.New(err)
	}

	log.Println(groups)
	//TODO: custom claims

	ok, token, err := auth.GenerateToken(username)
	if !ok {
		return &response, errors.New(err)
	}

	response.Token = token.TokenString
	response.TokenExpiration, _ = ptypes.TimestampProto(token.ExpirationTime)
	return &response, nil
}
