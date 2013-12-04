package services

import (
	"bones/repositories"
	"bones/validation"
	"bones/web/forms"
	"bones/web/sessions"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

const (
	serverError string = "Internal server error"
)

type TemplateRenderer interface {
	RenderTemplate(wr io.Writer, ctx TemplateContext) error
}

type Shortcuts struct {
	TemplateRenderer
	SessionStore sessions.SessionStore
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
	return &BaseContext{TemplateName: templateName, CsrfToken: csrfToken}
}

type TemplateContext interface {
	// Add an error that should be displayed
	// to the user
	AddError(err error)

	// Add a notice/message that should be
	// displayed to the user
	AddNotice(notice string)

	// Name of the (main) template that is being
	// rendered (not header or footer templates).
	Name() string
}

type BaseContext struct {
	TemplateName string
	Errors       []string
	Notices      []string
	CsrfToken    string
}

func (ctx *BaseContext) AddError(err error) {
	switch t := err.(type) {
	case *validation.ValidationError:
		for _, message := range t.Messages {
			ctx.Errors = append(ctx.Errors, message)
		}
	default:
		ctx.Errors = append(ctx.Errors, t.Error())
	}
}

func (ctx *BaseContext) AddNotice(notice string) {
	ctx.Notices = append(ctx.Notices, notice)
}

func (ctx *BaseContext) Name() string {
	return ctx.TemplateName
}

func (ctx *BaseContext) CsrfTokenField() template.HTML {
	inputField := fmt.Sprintf(`<input type="hidden" id="crsf-token" name="CsrfToken" value="%s">`, ctx.CsrfToken)
	return template.HTML(inputField)
}
