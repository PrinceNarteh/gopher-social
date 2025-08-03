package utils

import (
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	Logger.Errorw("internale server error:", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	if err := writeJSON(w, http.StatusInternalServerError, responseType{
		Status: "error",
		Error: &errorResponse{
			Code: http.StatusInternalServerError,
			Msg:  "the server encountered a problem",
		},
	}); err != nil {
		Logger.Errorw("internale server error:", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	}
}
