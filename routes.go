package main

import (
	"gogo-gadget-weather/weather"

	"github.com/go-chi/chi/v5"
)

func addRoutes(r *chi.Mux, w weather.Weather) {
	r.Get("/api/weather", w.HandleWeatherGet)
}
