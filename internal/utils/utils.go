package utils

import (
	"crypto/sha256"
	"encoding/hex"
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

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(hash[:])
	return hashedToken
}
