package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "Ok",
		"env":     app.config.Env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
	}
}
