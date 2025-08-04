package handlers

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"

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

func (h *UserHandler) ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	if err := h.store.User.Activate(r.Context(), token); err != nil {
		switch err {
		case store.ErrNotFound:
			utils.WriteError(w, http.StatusBadRequest, err)
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}
	utils.WriteResponse(w, r, http.StatusNoContent, "")
}

// GetUserHandler godoc
// @Summary      Get user details
// @Description  Retrieve user details by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        postId  path      int64  true  "User ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /users/{postId} [get]
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
