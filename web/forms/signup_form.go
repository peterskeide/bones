package forms

import (
	"bones/entities"
	"bones/repositories"
	"bones/validation"
	"errors"
)

type SignupForm struct {
	Users                repositories.UserRepository `schema:"_"`
	Email                string                      `schema:"email"`
	EmailConfirmation    string                      `schema:"email-confirmation"`
	Password             string                      `schema:"password"`
	PasswordConfirmation string                      `schema:"password-confirmation"`
}

func (f *SignupForm) Validate() error {
	validate := validation.New()

	validate.String(f.Email).NotEmpty("Email cannot be blank").Equals(f.EmailConfirmation, "Email didn't match email confirmation")
	validate.String(f.Password).NotEmpty("Password cannot be blank").Equals(f.PasswordConfirmation, "Password didn't match password confirmation")

	return validate.Result()
}

func (f *SignupForm) Save() error {
	_, err := f.Users.FindByEmail(f.Email)

	if err != nil {
		if err == repositories.NotFoundError {
			return f.saveNewUser()
		}

		return err
	}

	return errors.New("A user with that email address is already registered")
}

func (f *SignupForm) saveNewUser() error {
	user := entities.User{Email: f.Email}

	err := user.SetPassword(f.Password)

	if err != nil {
		return err
	}

	return f.Users.Insert(&user)
}
