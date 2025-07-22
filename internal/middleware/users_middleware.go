package middleware

import (
	"context"
	"net/http"

	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type userKey string

const UserCtx userKey = "user"

func (s *Middleware) UserContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.GetURLParamAsInt(r, "userId")
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, "invalid user id")
			return
		}

		ctx := r.Context()

		user, err := s.Store.User.FindById(ctx, userId)
		if err != nil {
			utils.WriteError(w, http.StatusNotFound, store.ErrNotFound.Error())
			return
		}

		ctx = context.WithValue(ctx, UserCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
