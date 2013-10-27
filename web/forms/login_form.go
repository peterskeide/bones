package forms

import (
	"bones/entities"
	"bones/repositories"
	"bones/validation"
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
	"log"
	"net/http"
)

var LoginFailedError = errors.New("Login failed")

type LoginForm struct {
	Request  *http.Request  `schema:"-"`
	User     *entities.User `schema:"-"`
	Email    string         `schema:"email"`
	Password string         `schema:"password"`
}

func (f *LoginForm) Validate() error {
	validate := validation.New()

	validate.String(f.Email).NotEmpty("Email cannot be blank")
	validate.String(f.Password).NotEmpty("Password cannot be blank")

	return validate.Result()
}

func (f *LoginForm) Save() error {
	var err error

	f.User, err = repositories.Users.FindByEmail(f.Email)

	if err != nil {
		if err == repositories.NotFoundError {
			log.Println("User not found")
			return LoginFailedError
		}

		log.Println("Repository error:", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(f.User.Password), []byte(f.Password))

	if err != nil {
		log.Println("Passwords didn't match")
		return err
	}

	// TODO save session to repository, update cookie in action

	return nil
}
