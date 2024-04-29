package weather

import (
	"fmt"
	"gogo-gadget-weather/customerror"
	"gogo-gadget-weather/openweathermap"
	"gogo-gadget-weather/serialize"
	"log"
	"net/http"
	"strconv"
)

const (
	mustBeNumericMessage  = "must be numeric value"
	isRequiredMessage     = "is required"
	latitudeRangeMessage  = "latitude must be between -90.0 and 90.0"
	longitudeRangeMessage = "longitude must be between -180.0 and 180"
	defaultCondition      = "Can't determine condition"
)

type WeatherResponse struct {
	Weather struct {
		Condition   string `json:"condition"`
		Temperature string `json:"temperature"`
	} `json:"weather"`
}

type weather struct {
	OpenWeatherMap openweathermap.OpenWeather
}

func New(ow openweathermap.OpenWeather) *weather {
	return &weather{
		OpenWeatherMap: ow,
	}
}

type Weather interface {
	HandleWeatherGet(w http.ResponseWriter, r *http.Request)
}

func (wtr weather) HandleWeatherGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lat := r.URL.Query().Get("lat")
	if err := validateLatitude(lat); err != nil {
		customerror.Api(
			w,
			err.Error(),
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		)
		return
	}

	long := r.URL.Query().Get("long")
	if err := validateLongitude(long); err != nil {
		customerror.Api(
			w,
			err.Error(),
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")
	response, err := wtr.OpenWeatherMap.FetchWeather(ctx, lat, long, apiKey)
	if err != nil {
		// In a production environment it would be unwise to return back this error so log it
		// and return a generic message
		log.Println("Error when fetching weather: " + err.Error())
		message := "Error processing your request. Please try again or open a support ticket"
		customerror.Api(
			w,
			message,
			http.StatusUnprocessableEntity,
			http.StatusText(http.StatusUnprocessableEntity),
		)
		return
	}

	if response.Code == http.StatusOK {
		err := serialize.Encode(w, http.StatusOK, formatOpenWeatherResponse(*response))
		if err != nil {
			log.Println("Error encoding openweather response: " + err.Error())
		}
		return
	}

	customerror.Api(w, response.Message, response.Code, http.StatusText(response.Code))
}

func validateLatitude(value string) error {
	if value == "" {
		return fmt.Errorf("latitude %s", isRequiredMessage)
	}

	lat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("latitude %s", mustBeNumericMessage)
	}

	if lat < -90 || lat > 90 {
		return fmt.Errorf(latitudeRangeMessage)
	}

	return nil
}

func validateLongitude(value string) error {
	if value == "" {
		return fmt.Errorf("longitude %s", isRequiredMessage)
	}

	long, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("longitude %s", mustBeNumericMessage)
	}

	if long < -180 || long > 180 {
		return fmt.Errorf(longitudeRangeMessage)
	}

	return nil
}

func formatOpenWeatherResponse(r openweathermap.ApiResponse) WeatherResponse {
	var weatherRepsonse WeatherResponse
	condition := defaultCondition
	if len(r.Weather) > 0 {
		condition = r.Weather[0].Description
	}

	weatherRepsonse.Weather.Temperature = convertTemperatureToFeeling(r.Main.FeelsLike)
	weatherRepsonse.Weather.Condition = condition

	return weatherRepsonse
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
