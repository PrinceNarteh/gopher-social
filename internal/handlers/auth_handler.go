package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/PrinceNarteh/social/internal/config"
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

// LoginHandler godoc
// @Summary      Log in user
// @Description  Log in user with email and password
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        loginDto  body      LoginDto  true  "Login DTO"
// @Success      200 {object}  models.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /auth/login [post]
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

// RegisterHandler godoc
// @Summary      Register user
// @Description  Register a new user with email and password
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User DTO"
// @Success      201 {object}  models.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /auth/register [post]
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

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	user.Password = hashedPassword

	if err := h.store.User.CreateAndInvite(ctx, &user, "uuid", config.Envs.MailConfig.Exp); err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			utils.WriteError(w, http.StatusBadRequest, err)
		case store.ErrDuplicateEmail:
			utils.WriteError(w, http.StatusBadRequest, err)
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}

	utils.WriteResponse(w, r, http.StatusCreated, user)
}
