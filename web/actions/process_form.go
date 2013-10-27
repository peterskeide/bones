package actions

import (
	"bones/web/forms"
	"net/http"
)

func ProcessForm(req *http.Request, form forms.Form) error {
	if err := forms.DecodeForm(form, req); err != nil {
		return err
	}

	if err := form.Validate(); err != nil {
		return err
	}

	return form.Save()
}
