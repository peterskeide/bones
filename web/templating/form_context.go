package templating

import (
	"bones/repositories"
	"encoding/hex"
	"fmt"
	"github.com/gorilla/securecookie"
	"html/template"
	"net/http"
)

type FormContext struct {
	*BaseContext
	CsrfToken string
}

func NewFormContext(res http.ResponseWriter, req *http.Request, base *BaseContext) (*FormContext, error) {
	csrfToken, err := getOrCreateCsrfToken(res, req)

	if err != nil {
		return nil, err
	}

	return &FormContext{base, csrfToken}, nil
}

func (fc *FormContext) CsrfTokenField() template.HTML {
	inputField := fmt.Sprintf(`<input type="hidden" id="crsf-token" name="CsrfToken" value="%s">`, fc.CsrfToken)
	return template.HTML(inputField)
}

func getOrCreateCsrfToken(res http.ResponseWriter, req *http.Request) (string, error) {
	session := repositories.Session(res, req)

	token, ok := session.Value("CsrfToken").(string)

	if !ok {
		randomKey := securecookie.GenerateRandomKey(32)
		token = hex.EncodeToString(randomKey)
		session.SetValue("CsrfToken", token)

		err := session.Save()

		if err != nil {
			return "", err
		}
	}

	return token, nil
}
