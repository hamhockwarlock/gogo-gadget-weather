package main

import (
	"fmt"
)

const openWeatherApiUrl = "https://api.openweathermap.org/data/2.5/weather"

type OpenWeatherWeatherResponse struct {
	Description string `json:"description"`
}

type OpenWeatherMainResponse struct {
	FeelsLike float64 `json:"feels_like"`
}

type OpenWeatherApiResponse struct {
	Weather []*OpenWeatherWeatherResponse `json:"weather"`
	Main    *OpenWeatherMainResponse      `json:"main"`
	Code    int                           `json:"cod"`
	Message string                        `json:"message"`
}

func fetchWeather(apiKey, lat, long string) (*OpenWeatherApiResponse, error) {
	url := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=imperial", openWeatherApiUrl, lat, long, apiKey)

	return get[OpenWeatherApiResponse](url)
}
