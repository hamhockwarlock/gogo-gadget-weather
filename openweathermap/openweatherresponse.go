package openweathermap

type WeatherResponse struct {
	Description string `json:"description"`
}

type MainResponse struct {
	FeelsLike float64 `json:"feels_like"`
}

type ApiResponse struct {
	Weather []*WeatherResponse `json:"weather"`
	Main    *MainResponse      `json:"main"`
	Code    int                `json:"cod"`
	Message string             `json:"message"`
}
