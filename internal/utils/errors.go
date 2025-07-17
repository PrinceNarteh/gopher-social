package utils

import (
	"fmt"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) error {
	fmt.Printf("internale server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	return writeJSON(w, http.StatusInternalServerError, responseType{
		Status: "error",
		Error: &errorResponse{
			Code: http.StatusInternalServerError,
			Msg:  "the server encountered a problem",
		},
	})
}
