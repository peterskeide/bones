package handlers

import (
	"bones/repositories"
	"bones/web/forms"
	"errors"
	"log"
	"net/http"
)

const (
	serverError string = "Internal server error"
)

type Shortcuts struct {
	TemplateRenderer
	SessionStore SessionStore
}

func (s Shortcuts) RenderPage(res http.ResponseWriter, pageContext TemplateContext) {
	err := s.RenderTemplate(res, pageContext)

	if err != nil {
		log.Println("Error when rendering template:", err, ". Context:", pageContext)
		http.Error(res, serverError, http.StatusInternalServerError)
	}
}

func (s Shortcuts) RenderPageWithErrors(res http.ResponseWriter, pageContext TemplateContext, errors ...error) {
	for _, err := range errors {
		pageContext.AddError(err)
	}

	s.RenderPage(res, pageContext)
}

func (s Shortcuts) Render404(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)

	err := s.RenderTemplate(res, s.TemplateContext(res, req, "404.html"))

	if err != nil {
		log.Println("Error when rendering 404 template:", err)
	}
}

func (s Shortcuts) Render401(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusUnauthorized)

	err := s.RenderTemplate(res, s.TemplateContext(res, req, "401.html"))

	if err != nil {
		log.Println("Error when rendering 401 template:", err)
	}
}

func (s Shortcuts) Render500(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusInternalServerError)

	err := s.RenderTemplate(res, s.TemplateContext(res, req, "500.html"))

	if err != nil {
		log.Println("Error when rendering 500 template:", err)
	}
}

func (s Shortcuts) FindEntityOr404(res http.ResponseWriter, req *http.Request, ef repositories.EntityFinder, id int) interface{} {
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

func (s Shortcuts) RedirectToLogin(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/login", http.StatusFound)
}

func (s Shortcuts) DecodeAndValidate(req *http.Request, form forms.Form) error {
	err := forms.DecodeForm(form, req)

	if err != nil {
		return err
	}

	return form.Validate()
}

func (s Shortcuts) TemplateContext(res http.ResponseWriter, req *http.Request, templateName string) *BaseContext {
	session := s.SessionStore.Session(res, req)
	csrfToken := session.CsrfToken()

	ctx := &BaseContext{TemplateName: templateName, CsrfToken: csrfToken}

	flashNotices := session.FlashNotices()

	for _, notice := range flashNotices {
		ctx.AddNotice(notice)
	}

	flashErrors := session.FlashErrors()

	for _, err := range flashErrors {
		ctx.AddError(errors.New(err))
	}

	err := session.Save()

	if err != nil {
		log.Println("Error saving session when creating new TemplateContext:", err)
	}

	return ctx
}

func (s Shortcuts) AddFlashError(res http.ResponseWriter, req *http.Request, msg string) error {
	session := s.SessionStore.Session(res, req)
	session.AddFlashError(msg)
	return session.Save()
}

func (s Shortcuts) AddFlashNotice(res http.ResponseWriter, req *http.Request, msg string) error {
	session := s.SessionStore.Session(res, req)
	session.AddFlashNotice(msg)
	return session.Save()
}
