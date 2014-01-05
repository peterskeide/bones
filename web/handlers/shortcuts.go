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

func (s Shortcuts) RenderPage(res http.ResponseWriter, req *http.Request, pageContext TemplateContext) {
	session := s.SessionStore.Session(res, req)
	err := session.Save()

	if err != nil {
		log.Println("Error saving session before rendering page:", err)
	}

	err = s.RenderTemplate(res, pageContext)

	if err != nil {
		log.Println("Error when rendering template:", err, ". Context:", pageContext)
	}
}

func (s Shortcuts) RenderPageWithErrors(res http.ResponseWriter, req *http.Request, pageContext TemplateContext, errors ...error) {
	for _, err := range errors {
		pageContext.AddError(err)
	}

	s.RenderPage(res, req, pageContext)
}

func (s Shortcuts) Render404(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	s.RenderPage(res, req, s.TemplateContext(res, req, "404.html"))
}

func (s Shortcuts) Render401(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusUnauthorized)
	s.RenderPage(res, req, s.TemplateContext(res, req, "401.html"))
}

func (s Shortcuts) Render500(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusInternalServerError)
	s.RenderPage(res, req, s.TemplateContext(res, req, "500.html"))
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
	s.redirect(res, req, "/login")
}

func (s Shortcuts) redirect(res http.ResponseWriter, req *http.Request, location string) {
	session := s.SessionStore.Session(res, req)
	err := session.Save()

	if err != nil {
		log.Println("Error saving session before redirecting to", location)
	}

	http.Redirect(res, req, location, http.StatusFound)
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

	return ctx
}

func (s Shortcuts) AddFlashError(res http.ResponseWriter, req *http.Request, msg string) {
	session := s.SessionStore.Session(res, req)
	session.AddFlashError(msg)
}

func (s Shortcuts) AddFlashNotice(res http.ResponseWriter, req *http.Request, msg string) {
	session := s.SessionStore.Session(res, req)
	session.AddFlashNotice(msg)
}
