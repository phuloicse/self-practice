package service

import (
	"errors"
	"jwt-app/model"
)

var mockUser = model.User{
	Username: "loiluong",
	Password: "1234",
}

func Authenticate(username, password string) error {
	if username == mockUser.Username && password == mockUser.Password {
		return nil
	}
	return errors.New("invalid credentials")
}
