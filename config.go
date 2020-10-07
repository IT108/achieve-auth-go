package main

import "os"

var tarantoolUser = "internal"
var tarantoolHost = "127.0.0.1:3301"
var tarantoolPassword = ""
var tarantooolAuthSpace = "auth"

func getConfig() {
	set(&tarantoolUser, os.Getenv("tarantool_user"))
	set(&tarantoolPassword, os.Getenv("tarantool_password"))
	set(&tarantoolHost, os.Getenv("tarantool_host"))
}

func set(variable *string, data string) {
	if data != "" {
		*variable = data
	}
}
