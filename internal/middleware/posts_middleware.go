package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type postKey string

const postCtx postKey = "post"

func (s *Middleware) PostContextMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postId, err := utils.GetURLParamAsInt(r, "postId")
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, "invalid post ID")
			return
		}

		ctx := r.Context()
		post, err := s.Store.Post.GetByID(ctx, postId)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				utils.WriteError(w, http.StatusNotFound, err.Error())
			default:
				utils.InternalServerError(w, r, err)
			}
			return
		}
		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetPostFromCtx(r *http.Request) *models.Post {
	if post, ok := r.Context().Value(postCtx).(*models.Post); ok {
		return post
	}
	return nil
}
