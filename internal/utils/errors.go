package utils

import (
	"fmt"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("internale server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	if err := writeJSON(w, http.StatusInternalServerError, responseType{
		Status: "error",
		Error: &errorResponse{
			Code: http.StatusInternalServerError,
			Msg:  "the server encountered a problem",
		},
	}); err != nil {
		fmt.Printf("internale server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	}
}
