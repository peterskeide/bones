package handlers

import (
	"io"
	"net/http"
)

type TemplateRenderer interface {
	RenderTemplate(wr io.Writer, ctx TemplateContext) error
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

type SessionStore interface {
	Session(res http.ResponseWriter, req *http.Request) Session
}

type Session interface {
	SetUserId(id int)
	UserId() int
	CsrfToken() string
	Clear() error
	Save() error
}
