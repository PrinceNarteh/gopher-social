package handlers

import (
	"github.com/PrinceNarteh/social/internal/middleware"
	"github.com/PrinceNarteh/social/internal/store"
)

type Handler struct {
	Auth       *AuthHandler
	Post       *PostHandler
	User       *UserHandler
	Feed       *FeedHandler
	Middleware *middleware.Middleware
}

func NewHandlers(store *store.Storage) *Handler {
	return &Handler{
		Auth:       &AuthHandler{store: store},
		Post:       &PostHandler{store: store},
		User:       &UserHandler{store: store},
		Middleware: &middleware.Middleware{Store: store},
	}
}
