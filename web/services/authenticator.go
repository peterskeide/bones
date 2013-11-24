package services

import (
	"bones/entities"
	"bones/repositories"
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
)

var LoginFailedError = errors.New("Login failed")

type EmailAuthenticator struct {
	Users repositories.UserRepository
}

func (auth EmailAuthenticator) Authenticate(email string, password string) (*entities.User, error) {
	user, err := auth.Users.FindByEmail(email)

	if err != nil {
		if err == repositories.NotFoundError {
			return nil, LoginFailedError
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, LoginFailedError
	}

	return user, nil
}
