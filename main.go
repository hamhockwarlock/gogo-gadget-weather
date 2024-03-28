package main

import (
	"gogo-gadget-weather/openweathermap"
	"gogo-gadget-weather/weather"
	"log"
	"net/http"
	"time"
)

const openWeatherApiUrl = "https://api.openweathermap.org/data/2.5/weather"

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running http server: %s\n", err)
	}
}

func run() error {
	httpClient := http.Client{}
	openWeatherMap := openweathermap.New(openWeatherApiUrl, httpClient, 30*time.Second)
	weather := weather.New(openWeatherMap)

	r := router(weather)
	err := http.ListenAndServe(":4242", r)
	if err != nil {
		return err
	}

	return nil
}
