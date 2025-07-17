package utils

import (
	"fmt"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi/v5"
)

func GetURLParamAsInt(r *http.Request, paramName string) (int64, error) {
	paramId := chi.URLParam(r, paramName)

	postId, err := strconv.ParseInt(paramId, 10, 64)
	fmt.Println(err)
	if err != nil {
		return 0, err
	}

	return postId, nil
}
