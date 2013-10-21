package forms

import (
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

func DecodeForm(form interface{}, req *http.Request) error {
	err := req.ParseForm()

	if err != nil {
		return err
	}

	return decoder.Decode(form, req.PostForm)
}
