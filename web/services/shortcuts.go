package services

import (
	"bones/repositories"
	"bones/web/forms"
	"bones/web/sessions"
	"bones/web/templating"
	"log"
	"net/http"
)

const (
	serverError string = "Internal server error"
)

type Shortcuts interface {
	RenderPage(res http.ResponseWriter, pageContext templating.TemplateContext)
	RenderPageWithErrors(res http.ResponseWriter, pageContext templating.TemplateContext, errors ...error)
	Render404(res http.ResponseWriter, req *http.Request)
	Render401(res http.ResponseWriter, req *http.Request)
	Render500(res http.ResponseWriter, req *http.Request)
	FindEntityOr404(res http.ResponseWriter, req *http.Request, ef repositories.EntityFinder, id int) interface{}
	RedirectToLogin(res http.ResponseWriter, req *http.Request)
	DecodeAndValidate(req *http.Request, form forms.Form) error
	TemplateContext(res http.ResponseWriter, req *http.Request, templateName string) *templating.BaseContext
}

type TemplatingShortcuts struct {
	templating.TemplateRenderer
	SessionStore sessions.SessionStore
}

func (s TemplatingShortcuts) RenderPage(res http.ResponseWriter, pageContext templating.TemplateContext) {
	err := s.RenderTemplate(res, pageContext)

	if err != nil {
		log.Println("Error when rendering template:", err, ". Context:", pageContext)
		http.Error(res, serverError, http.StatusInternalServerError)
	}
}

func (s TemplatingShortcuts) RenderPageWithErrors(res http.ResponseWriter, pageContext templating.TemplateContext, errors ...error) {
	for _, err := range errors {
		pageContext.AddError(err)
	}

	s.RenderPage(res, pageContext)
}

func (s TemplatingShortcuts) Render404(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)

	err := s.RenderTemplate(res, s.TemplateContext(res, req, "404.html"))

	if err != nil {
		log.Println("Error when rendering 404 template:", err)
	}
}

func (s TemplatingShortcuts) Render401(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusUnauthorized)

	err := s.RenderTemplate(res, s.TemplateContext(res, req, "401.html"))

	if err != nil {
		log.Println("Error when rendering 401 template:", err)
	}
}

func (s TemplatingShortcuts) Render500(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusInternalServerError)

	err := s.RenderTemplate(res, s.TemplateContext(res, req, "500.html"))

	if err != nil {
		log.Println("Error when rendering 500 template:", err)
	}
}

func (s TemplatingShortcuts) FindEntityOr404(res http.ResponseWriter, req *http.Request, ef repositories.EntityFinder, id int) interface{} {
	entity, err := ef.Find(id)

	if err != nil {
		if err == repositories.NotFoundError {
			s.Render404(res, req)
		} else {
			log.Println("Error on entity find:", err)
			s.Render500(res, req)
		}

		return nil
	}

	return entity
}

func (s TemplatingShortcuts) RedirectToLogin(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/login", http.StatusFound)
}

func (s TemplatingShortcuts) DecodeAndValidate(req *http.Request, form forms.Form) error {
	err := forms.DecodeForm(form, req)

	if err != nil {
		return err
	}

	return form.Validate()
}

func (s TemplatingShortcuts) TemplateContext(res http.ResponseWriter, req *http.Request, templateName string) *templating.BaseContext {
	session := s.SessionStore.Session(res, req)
	csrfToken := session.CsrfToken()
	return &templating.BaseContext{TemplateName: templateName, CsrfToken: csrfToken}
}
