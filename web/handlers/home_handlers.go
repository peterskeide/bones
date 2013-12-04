package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/services"
	"errors"
	"log"
	"net/http"
)

type HomeContext struct {
	*services.BaseContext
	Users []entities.User
}

type HomeHandler struct {
	services.Shortcuts
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

	h.RenderPage(res, &ctx)
}
