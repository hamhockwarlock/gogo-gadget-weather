package main

import (
	"fmt"
	"testing"
)

const (
	valueThatCannotBeCoerced = "hi!"
	valueOutsideRange        = "1000"
)

func TestLatitudeThatIsUncoercable(t *testing.T) {
	expectedError := fmt.Sprintf("latitude %s", mustBeNumericMessage)

	if err := validateLatitude(valueThatCannotBeCoerced); err.Error() != expectedError {
		t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
	}
}

func TestLongitudeThatIsUncoercable(t *testing.T) {
	expectedError := fmt.Sprintf("longitude %s", mustBeNumericMessage)

	if err := validateLongitude(valueThatCannotBeCoerced); err.Error() != expectedError {
		t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
	}
}

func TestLatitudeThatIsEmpty(t *testing.T) {
	expectedError := fmt.Sprintf("latitude %s", isRequiredMessage)

	if err := validateLatitude(""); err.Error() != expectedError {
		t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
	}
}

func TestLongitudeThatIsEmpty(t *testing.T) {
	expectedError := fmt.Sprintf("longitude %s", isRequiredMessage)

	if err := validateLongitude(""); err.Error() != expectedError {
		t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
	}
}

func TestLongitudeOutsideRange(t *testing.T) {
	if err := validateLongitude(valueOutsideRange); err.Error() != longitudeRangeMessage {
		t.Errorf("returned the wrong message: got %s want %s", err.Error(), longitudeRangeMessage)
	}
}

func TestLatitudeOutsideRange(t *testing.T) {
	if err := validateLatitude(valueOutsideRange); err.Error() != latitudeRangeMessage {
		t.Errorf("returned the wrong message: got %s want %s", err.Error(), latitudeRangeMessage)
	}
}

func TestLongitudeWithinRange(t *testing.T) {
	if err := validateLongitude("179"); err != nil {
		t.Errorf("returned the wrong message: got %s want nil", err.Error())
	}
}

func TestLatitudeWithinRange(t *testing.T) {
	if err := validateLatitude("89"); err != nil {
		t.Errorf("returned the wrong message: got %s want nil", err.Error())
	}
}

func TestConvertTemperatureToFeelingLessThan32(t *testing.T) {
	temperature := 11.0

	if feeling := convertTemperatureToFeeling(temperature); feeling != "cold" {
		t.Errorf("returned the wrong feeling: got %s want %s", feeling, "cold")
	}
}

func TestConvertTemperatureToFeelingGreater32LessThan76(t *testing.T) {
	temperature := 75.0

	if feeling := convertTemperatureToFeeling(temperature); feeling != "moderate" {
		t.Errorf("returned the wrong feeling: got %s want %s", feeling, "moderate")
	}
}

func TestConvertTemperatureToFeelingGreaterThan76(t *testing.T) {
	temperature := 100.0

	if feeling := convertTemperatureToFeeling(temperature); feeling != "hot" {
		t.Errorf("returned the wrong feeling: got %s want %s", feeling, "hot")
	}
}

func TestFormatOpenWeatherResponseHappyCase(t *testing.T) {
	weather := &OpenWeatherWeatherResponse{
		Description: "Snowy",
	}

	main := &OpenWeatherMainResponse{FeelsLike: 75}

	openWeatherResponse := OpenWeatherApiResponse{
		Main:    main,
		Code:    200,
		Weather: []*OpenWeatherWeatherResponse{weather},
	}

	formattedWeatherResponse := formatOpenWeatherResponse(openWeatherResponse)
	if formattedWeatherResponse.Weather.Condition != weather.Description {
		t.Errorf(
			"returned the wrong condition: got %s want %s",
			formattedWeatherResponse.Weather.Condition,
			weather.Description,
		)
	}

	if formattedWeatherResponse.Weather.Temperature != "moderate" {
		t.Errorf(
			"returned the wrong temperature: got %s want %s",
			formattedWeatherResponse.Weather.Temperature,
			"moderate",
		)
	}
}

func TestFormatOpenWeatherResponseMultipleWeathersPicksFirst(t *testing.T) {
	weather1 := &OpenWeatherWeatherResponse{
		Description: "Snowy",
	}

	weather2 := &OpenWeatherWeatherResponse{
		Description: "Rainy",
	}

	main := &OpenWeatherMainResponse{FeelsLike: 75}

	openWeatherResponse := OpenWeatherApiResponse{
		Main:    main,
		Code:    200,
		Weather: []*OpenWeatherWeatherResponse{weather1, weather2},
	}

	formattedWeatherResponse := formatOpenWeatherResponse(openWeatherResponse)
	if formattedWeatherResponse.Weather.Condition != weather1.Description {
		t.Errorf(
			"returned the wrong condition: got %s want %s",
			formattedWeatherResponse.Weather.Condition,
			weather1.Description,
		)
	}
}

func TestFormatOpenWeatherResponseNoWeather(t *testing.T) {
	main := &OpenWeatherMainResponse{FeelsLike: 75}

	openWeatherResponse := OpenWeatherApiResponse{
		Main:    main,
		Code:    200,
		Weather: []*OpenWeatherWeatherResponse{},
	}

	formattedWeatherResponse := formatOpenWeatherResponse(openWeatherResponse)
	if formattedWeatherResponse.Weather.Condition != defaultCondition {
		t.Errorf(
			"returned the wrong condition: got %s want %s",
			formattedWeatherResponse.Weather.Condition,
			defaultCondition,
		)
	}
}
