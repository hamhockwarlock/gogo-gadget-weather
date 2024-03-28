package weather

import (
	"fmt"
	"gogo-gadget-weather/openweathermap"
	"gogo-gadget-weather/test"
	"testing"
)

const (
	valueThatCannotBeCoerced = "hi!"
	valueOutsideRange        = "1000"
)

func TestLatitudeThatIsUncoercable(t *testing.T) {
	expectedError := fmt.Sprintf("latitude %s", mustBeNumericMessage)

	if err := validateLatitude(valueThatCannotBeCoerced); err.Error() != expectedError {
		test.Fatal(t, err.Error(), expectedError)
	}
}

func TestLongitudeThatIsUncoercable(t *testing.T) {
	expectedError := fmt.Sprintf("longitude %s", mustBeNumericMessage)

	if err := validateLongitude(valueThatCannotBeCoerced); err.Error() != expectedError {
		test.Fatal(t, err.Error(), expectedError)
	}
}

func TestLatitudeThatIsEmpty(t *testing.T) {
	expectedError := fmt.Sprintf("latitude %s", isRequiredMessage)

	if err := validateLatitude(""); err.Error() != expectedError {
		test.Fatal(t, err.Error(), expectedError)
	}
}

func TestLongitudeThatIsEmpty(t *testing.T) {
	expectedError := fmt.Sprintf("longitude %s", isRequiredMessage)

	if err := validateLongitude(""); err.Error() != expectedError {
		test.Fatal(t, err.Error(), expectedError)
	}
}

func TestLongitudeOutsideRange(t *testing.T) {
	if err := validateLongitude(valueOutsideRange); err.Error() != longitudeRangeMessage {
		t.Fatal(t, err.Error(), longitudeRangeMessage)
	}
}

func TestLatitudeOutsideRange(t *testing.T) {
	if err := validateLatitude(valueOutsideRange); err.Error() != latitudeRangeMessage {
		test.Fatal(t, err.Error(), latitudeRangeMessage)
	}
}

func TestLongitudeWithinRange(t *testing.T) {
	if err := validateLongitude("179"); err != nil {
		test.Fatal(t, err.Error(), "nil")
	}
}

func TestLatitudeWithinRange(t *testing.T) {
	if err := validateLatitude("89"); err != nil {
		test.Fatal(t, err.Error(), "nil")
	}
}

func TestConvertTemperatureToFeelingLessThan32(t *testing.T) {
	temperature := 11.0

	if feeling := convertTemperatureToFeeling(temperature); feeling != "cold" {
		test.Fatal(t, feeling, "cold")
	}
}

func TestConvertTemperatureToFeelingGreater32LessThan76(t *testing.T) {
	temperature := 75.0

	if feeling := convertTemperatureToFeeling(temperature); feeling != "moderate" {
		test.Fatal(t, feeling, "moderate")
	}
}

func TestConvertTemperatureToFeelingGreaterThan76(t *testing.T) {
	temperature := 100.0

	if feeling := convertTemperatureToFeeling(temperature); feeling != "hot" {
		test.Fatal(t, feeling, "hot")
	}
}

func TestFormatOpenWeatherResponseHappyCase(t *testing.T) {
	weather := &openweathermap.WeatherResponse{
		Description: "Snowy",
	}

	main := &openweathermap.MainResponse{FeelsLike: 75}

	openWeatherResponse := openweathermap.ApiResponse{
		Main:    main,
		Code:    200,
		Weather: []*openweathermap.WeatherResponse{weather},
	}

	formattedWeatherResponse := formatOpenWeatherResponse(openWeatherResponse)
	if formattedWeatherResponse.Weather.Condition != weather.Description {
		test.Fatal(
			t,
			formattedWeatherResponse.Weather.Condition,
			weather.Description,
		)
	}

	if formattedWeatherResponse.Weather.Temperature != "moderate" {
		test.Fatal(
			t,
			formattedWeatherResponse.Weather.Temperature,
			"moderate",
		)
	}
}

func TestFormatOpenWeatherResponseMultipleWeathersPicksFirst(t *testing.T) {
	weather1 := &openweathermap.WeatherResponse{
		Description: "Snowy",
	}

	weather2 := &openweathermap.WeatherResponse{
		Description: "Rainy",
	}

	main := &openweathermap.MainResponse{FeelsLike: 75}

	openWeatherResponse := openweathermap.ApiResponse{
		Main:    main,
		Code:    200,
		Weather: []*openweathermap.WeatherResponse{weather1, weather2},
	}

	formattedWeatherResponse := formatOpenWeatherResponse(openWeatherResponse)
	if formattedWeatherResponse.Weather.Condition != weather1.Description {
		test.Fatal(t, formattedWeatherResponse.Weather.Condition, weather1.Description)
	}
}

func TestFormatOpenWeatherResponseNoWeather(t *testing.T) {
	main := &openweathermap.MainResponse{FeelsLike: 75}

	openWeatherResponse := openweathermap.ApiResponse{
		Main:    main,
		Code:    200,
		Weather: []*openweathermap.WeatherResponse{},
	}

	formattedWeatherResponse := formatOpenWeatherResponse(openWeatherResponse)
	if formattedWeatherResponse.Weather.Condition != defaultCondition {
		test.Fatal(t, formattedWeatherResponse.Weather.Condition, defaultCondition)
	}
}
