package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Weather struct {
	Condition   string `json:"condition"`
	Temperature string `json:"temperature"`
}
type WeatherResponse struct {
	Weather Weather `json:"weather"`
}

func handleWeatherGet(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	err := validateLatitude(lat)
	if err = validateLatitude(lat); err != nil {
		apiError(w, err.Error(), http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	long := r.URL.Query().Get("long")
	if err = validateLongitude(long); err != nil {
		apiError(w, err.Error(), http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	apiKey := r.Header.Get("X-API-KEY")
	response, err := fetchWeather(apiKey, lat, long)
	if err != nil {
		// In a production environment it would be unwise to return back this error so log it
		// and return a generic message
		log.Println("Error when fetching weather: " + err.Error())
		message := "Error processing your request. Please try again or open a support ticket"
		apiError(w, message, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))
		return
	}

	if response.Code == http.StatusOK {
		err := encode(w, http.StatusOK, formatOpenWeatherResponse(*response))
		if err != nil {
			log.Println("Error encoding openweather response: " + err.Error())
		}
		return
	}

	apiError(w, response.Message, response.Code, http.StatusText(response.Code))
}

func validateLatitude(value string) error {
	if value == "" {
		return fmt.Errorf("latitude is required")
	}

	lat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("latitude must be a numeric value")
	}

	if lat < -90 || lat > 90 {
		return fmt.Errorf("latitude must be between -90.0 and 90.0")
	}

	return nil
}

func validateLongitude(value string) error {
	if value == "" {
		return fmt.Errorf("longitude is required")
	}

	long, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("longitude must be a numeric value")
	}

	if long < -180 || long > 180 {
		return fmt.Errorf("longitude must be between -90.0 and 90.0")
	}

	return nil
}

func formatOpenWeatherResponse(r OpenWeatherApiResponse) WeatherResponse {
	condition := "Can't determine condition"
	if len(r.Weather) > 0 {
		condition = r.Weather[0].Description
	}

	return WeatherResponse{
		Weather: Weather{
			Temperature: convertTemperatureToFeeling(r.Main.FeelsLike),
			Condition:   condition,
		},
	}
}

func convertTemperatureToFeeling(temp float64) string {
	if temp <= 32 {
		return "cold"
	}
	if temp <= 75 {
		return "moderate"
	}

	return "hot"
}
