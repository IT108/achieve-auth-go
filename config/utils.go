package config

import "os"

func GetConfig() {
	set(&TarantoolUser, os.Getenv("tarantool_user"))
	set(&TarantoolPassword, os.Getenv("tarantool_password"))
	set(&TarantoolHost, os.Getenv("tarantool_host"))
}

func set(variable *string, data string) {
	if data != "" {
		*variable = data
	}
}
