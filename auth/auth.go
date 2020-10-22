package auth

import (
	. "github.com/IT108/achieve-auth-go/config"
	db "github.com/IT108/achieve-db-go"
	"log"
)

func Register(username string, email string, password string) (ok bool, err string) {
	ok, err = true, ""

	hashPass := hashAndSalt([]byte(password))
	result := db.Insert(TarantoolAuthSpace, []interface{}{email, username, hashPass, []string{}, false})

	if result.Code != 0 {
		ok = false
		emailAvail, _ := IsEmailAvailable(email)
		if !emailAvail {
			err = "email already registered"
		} else {
			err = "username already registered"
		}
	}

	return ok, err
}

func Authenticate(username string, password string) (ok bool, err string) {
	ok, err = true, ""

	query := *db.SelectUsers(TarantoolAuthSpace, "primary", username)
	if len(query) != 0 {
		if comparePasswords(query[0].PasswordHash, []byte(password)) {
			return ok, err
		}
	}

	query = *db.SelectUsers(TarantoolAuthSpace, "secondary", username)
	if len(query) != 0 {
		if comparePasswords(query[0].PasswordHash, []byte(password)) {
			return ok, err
		}
	}

	ok, err = false, "Incorrect password"
	return ok, err
}

func Authorize(email string) (ok bool, groups []string, err string) {
	ok, groups, err = true, []string{}, ""

	query := *db.SelectUsers(TarantoolAuthSpace, "primary", email)
	if len(query) == 0 {
		err := "Error on select user: no user with given username, " + email
		log.Println(err)

		return false, groups, err
	}

	groups = query[0].Groups
	return ok, groups, err
}
