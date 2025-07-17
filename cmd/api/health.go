package main

import (
	"net/http"

	"github.com/PrinceNarteh/social/internal/config"
	"github.com/PrinceNarteh/social/internal/utils"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "Ok",
		"env":     config.Envs.AppConfig.Env,
		"version": config.Envs.AppConfig.Version,
	}

	if err := utils.WriteResponse(w, http.StatusOK, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
