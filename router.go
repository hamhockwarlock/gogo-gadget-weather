package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func router() *chi.Mux {
	r := chi.NewRouter()

	setupMiddleware(r)
	addRoutes(r)
	r.NotFound(handleNotFound)
	r.MethodNotAllowed(handleMethodNotAllowed)

	return r
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("No endpoint at %s", r.RequestURI)
	apiError(w, message, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s not allowed for endpoint", r.Method)
	apiError(w, message, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
}
