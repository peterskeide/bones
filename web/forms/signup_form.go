package forms

import (
	"bones/entities"
	"bones/validation"
	"code.google.com/p/go.crypto/bcrypt"
)

type SignupForm struct {
	Email                string `schema:"email"`
	EmailConfirmation    string `schema:"email-confirmation"`
	Password             string `schema:"password"`
	PasswordConfirmation string `schema:"password-confirmation"`
}

func (f SignupForm) Validate() error {
	validate := validation.New()

	validate.String(f.Email).NotEmpty("Email cannot be blank").Equals(f.EmailConfirmation, "Email didn't match email confirmation")
	validate.String(f.Password).NotEmpty("Password cannot be blank").Equals(f.PasswordConfirmation, "Password didn't match password confirmation")

	return validate.Result()
}

func (f SignupForm) User() (*entities.User, error) {
	encryptedPassword, err := f.encryptPassword()

	if err != nil {
		return nil, err
	}

	return &entities.User{Email: f.Email, Password: encryptedPassword}, nil
}

func (f SignupForm) encryptPassword() (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(pwd), nil
}
