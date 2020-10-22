package main

import (
	db "github.com/IT108/achieve-db-go"
	models "github.com/IT108/achieve-models-go"
	"log"
	"net/http"
)

func register(request models.RegisterRequest) models.RegisterResponse {
	hashPass := hashAndSalt([]byte(request.Password))
	result := db.Insert(tarantoolAuthSpace, []interface{}{request.Email, request.Username, hashPass, []string{}, false})
	response := models.RegisterResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
	}
	if result.Code != 0 {
		response.ResponseCode = http.StatusConflict
		if isEmail(request.Email) {
			response.Error = "email already registered"
			return response
		} else {
			response.Error = "username already registered"
			return response
		}
	}
	return response
}

func isRegistered(request models.IsRegisteredRequest) models.IsRegisteredResponse {
	result := models.IsRegisteredResponse{
		Request:              request.Request,
		ResponseCode:         http.StatusOK,
		IsEmailRegistered:    false,
		IsUsernameRegistered: false,
		Error:                "",
	}
	if isEmail(request.Email) {
		result.IsEmailRegistered = true
		result.ResponseCode = http.StatusConflict
		result.Error = "email or username are already taken"
	}
	if isUsername(request.Username) {
		result.IsUsernameRegistered = true
		result.ResponseCode = http.StatusConflict
		result.Error = "email or username are already taken"
	}
	return result
}

func isEmailRegistered(request models.IsEmailRegisteredRequest) models.IsEmailResponse {
	result := models.IsEmailResponse{
		Request:           request.Request,
		ResponseCode:      http.StatusOK,
		IsEmailRegistered: false,
		Error:             "",
	}
	if isEmail(request.Email) {
		result.ResponseCode = http.StatusConflict
		result.Error = "email already taken"
		result.IsEmailRegistered = true
	}
	return result
}

func isUserRegistered(request models.IsUserRegisteredRequest) models.IsUserRegisteredResponse {
	result := models.IsUserRegisteredResponse{
		Request:              request.Request,
		ResponseCode:         http.StatusOK,
		IsUsernameRegistered: false,
		Error:                "",
	}
	if isUsername(request.Username) {
		result.ResponseCode = http.StatusConflict
		result.Error = "username already taken"
		result.IsUsernameRegistered = true
	}
	return result
}

func isEmail(email string) bool {
	result := db.Select(tarantoolAuthSpace, "primary", email).Data
	if len(result) != 0 {
		return true
	}
	return false
}

func isUsername(username string) bool {
	result := db.Select(tarantoolAuthSpace, "secondary", username).Data
	if len(result) != 0 {
		return true
	}
	return false
}

func authenticate(request models.AuthenticateRequest) models.AuthenticateResponse {
	result := models.AuthenticateResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
	}
	query := *db.SelectUsers(tarantoolAuthSpace, "primary", request.Username)
	if len(query) != 0 {
		if comparePasswords(query[0].PasswordHash, []byte(request.Password)) {
			return result
		}
	}
	query = *db.SelectUsers(tarantoolAuthSpace, "secondary", request.Username)

	if len(query) != 0 {
		if comparePasswords(query[0].PasswordHash, []byte(request.Password)) {
			return result
		}
	}
	result.ResponseCode = http.StatusForbidden
	result.Error = "Incorrect password"
	log.Println(result.ResponseCode)
	return result
}

func authorize(request models.AuthorizeRequest) models.AuthorizeResponse {
	query := *db.SelectUsers(tarantoolAuthSpace, "primary", request.User)
	result := models.AuthorizeResponse{
		Request:      request.Request,
		ResponseCode: http.StatusOK,
	}
	if len(query) == 0 {
		err := "Error on select user: no user with given username, " + request.User
		log.Println(err)
		result.ResponseCode = http.StatusConflict
		result.Error = err
		return result
	}
	result.Roles = query[0].Groups
	return result
}
