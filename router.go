package main

import (
	"fmt"
	"gogo-gadget-weather/customerror"
	"gogo-gadget-weather/weather"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func router(w weather.Weather) *chi.Mux {
	r := chi.NewRouter()

	setupMiddleware(r)
	addRoutes(r, w)
	r.NotFound(handleNotFound)
	r.MethodNotAllowed(handleMethodNotAllowed)

	return r
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("No endpoint at %s", r.RequestURI)
	customerror.Api(w, message, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s not allowed for endpoint", r.Method)
	customerror.Api(
		w,
		message,
		http.StatusMethodNotAllowed,
		http.StatusText(http.StatusMethodNotAllowed),
	)
}
