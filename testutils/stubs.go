package testutils

import (
	"bones/entities"
	"bones/web/sessions"
	"bones/web/templating"
	"io"
	"net/http"
)

type UserRepositoryStub struct {
}

func (stub UserRepositoryStub) Insert(user *entities.User) error {
	return nil
}

func (stub UserRepositoryStub) FindByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func (stub UserRepositoryStub) FindById(id int) (*entities.User, error) {
	return nil, nil
}

func (stub UserRepositoryStub) Find(id int) (interface{}, error) {
	return nil, nil
}

func (stub UserRepositoryStub) All() ([]entities.User, error) {
	return nil, nil
}

type SessionStub struct {
	StubbedUserId    int
	StubbedCsrfToken string
	ClearErr         error
	Cleared          bool
	Saved            bool
	SaveErr          error
}

func (stub *SessionStub) UserId() int {
	return stub.StubbedUserId
}

func (stub *SessionStub) SetUserId(id int) {
	stub.StubbedUserId = id
}

func (stub *SessionStub) CsrfToken() string {
	return stub.StubbedCsrfToken
}

func (stub *SessionStub) Clear() error {
	stub.Cleared = true
	return stub.ClearErr
}

func (stub *SessionStub) Save() error {
	stub.Saved = true
	return stub.SaveErr
}

type SessionStoreStub struct {
	CurrentSession *SessionStub
}

func (stub *SessionStoreStub) Session(res http.ResponseWriter, req *http.Request) sessions.Session {
	if stub.CurrentSession == nil {
		stub.CurrentSession = &SessionStub{}
	}

	return stub.CurrentSession
}

type TemplateRendererStub struct {
	Writer    io.Writer
	Ctx       templating.TemplateContext
	RenderErr error
}

func (stub *TemplateRendererStub) RenderTemplate(wr io.Writer, ctx templating.TemplateContext) error {
	stub.Writer = wr
	stub.Ctx = ctx
	return stub.RenderErr
}
