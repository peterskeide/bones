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
	ResponseWriter http.ResponseWriter `schema:"-"`
	Request        *http.Request       `schema:"-"`
	Users          repositories.UserRepository
	User           *entities.User `schema:"-"`
	// Need to include this because of gorilla/schema.
	// Schema should really ignore this field if it is
	// not declared in the struct or set to "-".
	// However, it currently returns an error on decode
	CsrfToken string `schema:"CsrfToken"`
	Email     string `schema:"email"`
	Password  string `schema:"password"`
}

func (f *LoginForm) Validate() error {
	validate := validation.New()

	validate.String(f.Email).NotEmpty("Email cannot be blank")
	validate.String(f.Password).NotEmpty("Password cannot be blank")

	return validate.Result()
}

func (f *LoginForm) Save() error {
	err := f.findAndAuthenticateUser()

	if err != nil {
		return err
	}

	return f.createSession()
}

func (f *LoginForm) findAndAuthenticateUser() error {
	var err error

	f.User, err = f.Users.FindByEmail(f.Email)

	if err != nil {
		if err == repositories.NotFoundError {
			return LoginFailedError
		}

		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(f.User.Password), []byte(f.Password))

	if err != nil {
		return LoginFailedError
	}

	return nil
}

func (f *LoginForm) createSession() error {
	session := repositories.Session(f.ResponseWriter, f.Request)
	session.SetValue("user_id", f.User.Id)
	err := session.Save()

	if err != nil {
		log.Println("Failed to save session:", err)
		return LoginFailedError
	}

	return nil
}
