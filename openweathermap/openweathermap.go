package openweathermap

import (
	"context"
	"fmt"
	"gogo-gadget-weather/rest"
	"net/http"
	"time"
)

type OpenWeather interface {
	FetchWeather(ctx context.Context, lat, long, apiKey string) (*ApiResponse, error)
}

func New(baseURL string, client http.Client, timeout time.Duration) *v1 {
	return &v1{
		baseURL: baseURL,
		client:  client,
		timeout: timeout,
	}
}

func (v *v1) FetchWeather(
	ctx context.Context,
	lat, long, apiKey string,
) (*ApiResponse, error) {
	url := fmt.Sprintf(
		"%s?lat=%s&lon=%s&appid=%s&units=imperial",
		v.baseURL,
		lat,
		long,
		apiKey,
	)

	return rest.Get[ApiResponse](&v.client, ctx, url, v.timeout)
}
