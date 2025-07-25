package handlers

import (
	"net/http"

	"github.com/PrinceNarteh/social/internal/middleware"
	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type UserHandler struct {
	store *store.Storage
}

func getUserFromCtx(r *http.Request) *models.User {
	if user, ok := r.Context().Value(middleware.UserCtx).(*models.User); ok {
		return user
	}
	return nil
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)
	utils.WriteResponse(w, r, http.StatusOK, user)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) FollowHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) UnfollowHandler(w http.ResponseWriter, r *http.Request) {}
