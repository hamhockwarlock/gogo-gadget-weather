package main

import (
	"fmt"
)

// In an ideal world this class would be moved to it's own package
// Where we could describe an interface, a struct designed for versioning making
// testing easier. Here's roughly what that might look like:

// type OpenWeather interface {
//   FetchWeather(ctx context.Context, lat, long string) (*OpenWeatherApiResponse, error)
// }

// type v1 struct {
//   baseURL string
//   client *http.Client
// }
// fetchWeather(apiKey, lat, long string) -> (v *v1) FetchWeather(ctx context.Context, lat, long string)

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
	url := fmt.Sprintf(
		"%s?lat=%s&lon=%s&appid=%s&units=imperial",
		openWeatherApiUrl,
		lat,
		long,
		apiKey,
	)

	return get[OpenWeatherApiResponse](url)
}
