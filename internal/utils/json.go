package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
}

type responseType struct {
	Status string         `json:"status"`
	Data   any            `json:"data,omitempty"`
	Error  *errorResponse `json:"error,omitempty"`
}

func ParseJSON(w http.ResponseWriter, r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("error")
	}

	maxBytesReader := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytesReader))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(payload)
}

func writeJSON(w http.ResponseWriter, statusCode int, payload responseType) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}

func WriteResponse(w http.ResponseWriter, r *http.Request, statusCode int, payload any) {
	if err := writeJSON(w, statusCode, responseType{
		Status: "success",
		Data:   payload,
	}); err != nil {
		InternalServerError(w, r, err)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, payload any) error {
	return writeJSON(w, statusCode, responseType{
		Status: "error",
		Error: &errorResponse{
			Code: statusCode,
			Msg:  payload,
		},
	})
}
