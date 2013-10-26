package entities

import (
	"code.google.com/p/go.crypto/bcrypt"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func (u *User) SetPassword(password string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(pwd)

	return nil
}
