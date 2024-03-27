package main

import (
	"github.com/go-chi/chi/v5"
)

func addRoutes(r *chi.Mux) {
  r.Get("/api/weather", WeatherHandler)
}
