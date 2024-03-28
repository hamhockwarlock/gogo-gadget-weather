package customerror

import (
	"gogo-gadget-weather/serialize"
	"log"
	"net/http"
)

type ApiErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func Api(w http.ResponseWriter, message string, status int, error string) {
	apiError := ApiErrorResponse{
		Message: message,
		Success: false,
		Error:   error,
	}

	err := serialize.Encode(w, status, apiError)
	if err != nil {
		log.Println("error encoding apiError: " + err.Error())
	}
}
