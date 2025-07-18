package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type LoginDto struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthHandler struct {
	store *store.Storage
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginDto LoginDto

	if err := utils.ParseJSON(w, r, &loginDto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateStruct(loginDto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusCreated, loginDto)
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := utils.ParseJSON(w, r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateStruct(user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()

	emailExists, err := h.store.User.FindByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if emailExists != nil {
		utils.WriteError(w, http.StatusConflict, "email already in exists")
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	user.Password = hashedPassword

	if err := h.store.User.Create(ctx, &user); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusCreated, user)
}
