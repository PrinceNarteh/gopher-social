package handlers

import "github.com/PrinceNarteh/social/internal/store"

type Handler struct {
	Post *PostHandler
	Auth *AuthHandler
}

func NewHandlers(store *store.Storage) *Handler {
	return &Handler{
		Post: &PostHandler{store: store},
		Auth: &AuthHandler{store: store},
	}
}
