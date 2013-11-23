package forms

import (
	"bones/validation"
)

type LoginForm struct {
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
