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

	utils.WriteResponse(w, r, http.StatusOK, data)
}
