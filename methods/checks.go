package methods

import (
	"context"
	"errors"
	"github.com/IT108/achieve-auth-go/auth"
	proto "github.com/IT108/achieve-auth-go/auth-proto"
	"log"
)


func (*Server) IsRegistered(ctx context.Context, request *proto.IsRegisteredRequest) (*proto.IsRegisteredResponse, error) {
	errorMessage := ""

	response := &proto.IsRegisteredResponse{
		IsUserRegistered:     false,
		IsEmailRegistered:    false,
	}

	email, err := auth.IsEmailAvailable(request.Email)
	if !email {
		response.IsEmailRegistered = true
		errorMessage = err
	}

	username, err := auth.IsUsernameAvailable(request.Email)
	if !username {
		response.IsUserRegistered = true
		errorMessage = err
	}

	if !username && !email {
		errorMessage = "Email and username are registered"
	}

	return response, errors.New(errorMessage)
}

func (s *Server) IsEmailRegistered(ctx context.Context, request *proto.IsEmailRequest) (*proto.IsEmailResponse, error) {
	log.Print("Is email")

	result := proto.IsEmailResponse{
		IsEmailRegistered: false,
	}

	ok, _ := auth.IsEmailAvailable(request.Email)
	if !ok {
		result.IsEmailRegistered = true
		return &result, errors.New("email already registered")
	}

	return &result, nil
}

func (*Server) IsUserRegistered(ctx context.Context, request *proto.IsUserRequest) (*proto.IsUserResponse, error) {
	result := &proto.IsUserResponse{
		IsUserRegistered: false,
	}

	ok, err := auth.IsUsernameAvailable(request.Username)
	if !ok {
		result.IsUserRegistered = true
		return result, errors.New(err)
	}

	return result, nil
}
