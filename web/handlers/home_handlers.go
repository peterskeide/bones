package handlers

import (
	"bones/entities"
	"bones/repositories"
	"errors"
	"log"
	"net/http"
)

type HomeContext struct {
	*BaseContext
	Users []entities.User
}

type HomeHandler struct {
	Shortcuts
	Users repositories.UserRepository
}

func (h *HomeHandler) LoadHomePage(res http.ResponseWriter, req *http.Request) {
	ctx := HomeContext{h.TemplateContext(res, req, "index.html"), nil}
	users, err := h.Users.All()

	if err != nil {
		log.Println("Error loading users from repository:", err)
		ctx.AddError(errors.New("Unable to load users"))
	} else {
		ctx.Users = users
	}

	h.RenderPage(res, req, &ctx)
}
