package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
  r := chi.NewRouter()

  setupMiddleware(r)
  addRoutes(r)

  return r
}

func main() {
  r := Router()
  http.ListenAndServe(":4242", r)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
  lat := r.URL.Query().Get("lat")
  long := r.URL.Query().Get("long")
  
  response := WeatherResponse{
    Latitude: lat,
    Longitude: long,
    Message: "Just testing",
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

type WeatherResponse struct {
  Latitude  string `json:"latitude"`
  Longitude string `json:"longitude"`
  Message   string `json:"message"`
}

func run() {}
