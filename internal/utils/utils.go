package utils

import (
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi/v5"
)

func GetURLParamAsInt(r *http.Request, paramName string) (int64, error) {
	paramId := chi.URLParam(r, paramName)

	postId, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		return 0, err
	}

	return postId, nil
}
