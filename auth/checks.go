package auth

import (
	. "github.com/IT108/achieve-auth-go/config"
	db "github.com/IT108/achieve-db-go"
)

func IsEmailAvailable(email string) (ok bool, err string) {
	ok, err = true, ""

	result := db.Select(TarantoolAuthSpace, "primary", email).Data
	if len(result) != 0 {
		ok, err = false, "email are already taken"
	}

	return ok, err
}

func IsUsernameAvailable(username string) (ok bool, err string) {
	ok, err = true, ""

	result := db.Select(TarantoolAuthSpace, "secondary", username).Data
	if len(result) != 0 {
		ok, err = false, "username are already taken"
	}

	return ok, err
}
