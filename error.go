package main

import (
	"encoding/json"
	"net/http"
)

type ApiErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"status"`
	Error   string `json:"error"`
}

// TODO: Move the encoding to a universal spot
func ApiError(w http.ResponseWriter, message string, status int, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiErrorResponse{Message: message, Success: false, Error: error})
}
