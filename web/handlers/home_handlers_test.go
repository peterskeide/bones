package handlers

import (
	"bones/entities"
	"bones/testutils"
	"bones/web/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userRepositoryStub struct {
	testutils.UserRepositoryStub
	users []entities.User
}

func (stub *userRepositoryStub) All() ([]entities.User, error) {
	return stub.users, nil
}

func TestLoadsHomePageRendersIndexTemplateWithAllUsers(t *testing.T) {
	templateRenderer := testutils.TemplateRendererStub{}
	sessionStore := testutils.SessionStoreStub{}
	shortcuts := services.TemplatingShortcuts{&templateRenderer, &sessionStore}

	users := []entities.User{
		entities.User{Id: 1, Email: "a@test.com"},
		entities.User{Id: 2, Email: "b@test.com"},
		entities.User{Id: 3, Email: "c@test.com"},
	}

	userRepository := userRepositoryStub{testutils.UserRepositoryStub{}, users}

	handler := HomeHandler{&shortcuts, &userRepository}

	responseWriter := httptest.NewRecorder()
	request := http.Request{}

	handler.LoadHomePage(responseWriter, &request)

	if templateRenderer.Ctx.Name() != "index.html" {
		t.Error("Expected HomeHandler to render index.html template")
	}

	if homeContext, ok := templateRenderer.Ctx.(*HomeContext); ok {
		if !entities.UserEquals(users, homeContext.Users) {
			t.Error("Expected HomeHandler to render a list of all users")
		}
	} else {
		t.Error("Expected HomeContext")
	}
}
