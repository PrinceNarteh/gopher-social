package handlers

import (
	"database/sql"
	"net/http"

	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type UserHandler struct {
	store *store.Storage
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetURLParamAsInt(r, "userId")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.store.User.FindById(r.Context(), userId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.WriteError(w, http.StatusNotFound, store.ErrNotFound.Error())
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}

	utils.WriteResponse(w, r, http.StatusOK, user)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
}
