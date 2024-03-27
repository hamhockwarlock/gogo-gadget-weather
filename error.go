package main

import (
	"log"
	"net/http"
)

type ApiErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"status"`
	Error   string `json:"error"`
}

func apiError(w http.ResponseWriter, message string, status int, error string) {
	apiError := ApiErrorResponse{
		Message: message,
		Success: false,
		Error:   error,
	}

	err := encode(w, status, apiError)
	if err != nil {
		log.Println("error encoding apiError: " + err.Error())
	}
}
