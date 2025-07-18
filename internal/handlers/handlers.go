package handlers

import (
	"github.com/PrinceNarteh/social/internal/middleware"
	"github.com/PrinceNarteh/social/internal/store"
)

type Handler struct {
	Post       *PostHandler
	Auth       *AuthHandler
	Middleware *middleware.Middleware
}

func NewHandlers(store *store.Storage) *Handler {
	return &Handler{
		Post:       &PostHandler{store: store},
		Auth:       &AuthHandler{store: store},
		Middleware: &middleware.Middleware{Store: store},
	}
}
