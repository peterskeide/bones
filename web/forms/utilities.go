package forms

import (
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

// Will panic if error
func DecodeForm(form interface{}, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		panic(err)
	}

	err = decoder.Decode(form, req.PostForm)

	if err != nil {
		panic(err)
	}
}

type query map[string]interface{}
