package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

const apiKeyHeaderName = "X-API-KEY"
const apiKeyRequiredMessage = "An OpenWeather API key must be provided"

func setupMiddleware(r *chi.Mux) {
	r.Use(apiKeyRequired)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
}

func apiKeyRequired(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(apiKeyHeaderName) == "" {
			apiError(w, apiKeyRequiredMessage, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		h.ServeHTTP(w, r)
	})
}
