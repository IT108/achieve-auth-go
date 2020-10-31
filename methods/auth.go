package methods

import (
	"context"
	"errors"
	"github.com/IT108/achieve-auth-go/auth"
	proto "github.com/IT108/achieve-auth-go/auth-proto"
	"github.com/golang/protobuf/ptypes"
	"log"
)

type Server struct {
}



func (s *Server) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.Response, error) {
	errorMessage := ""
	ok, err := auth.Register(request.Username, request.Email, request.Password)

	response := &proto.Response{
		Status: 200,
	}

	if !ok {
		response.Status = 500
		errorMessage = err
	}

	return response, errors.New(errorMessage)
}

func (s *Server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.Token, error) {
	ok, err := auth.Authenticate(request.Username, request.Password)

	response := proto.Token{
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
