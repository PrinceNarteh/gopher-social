package middleware

import (
	"github.com/PrinceNarteh/social/internal/store"
)

type Middleware struct {
	Store *store.Storage
}

func NewMiddleware(store *store.Storage) *Middleware {
	return &Middleware{
		Store: store,
	}
}
