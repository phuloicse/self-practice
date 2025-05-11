package service

import (
	"errors"
	config "jwt-app/config"
)

var mockUser = config.User

func Authenticate(username, password string) error {
	if username == mockUser.Username && password == mockUser.Password {
		return nil
	}
	return errors.New("invalid credentials")
}
