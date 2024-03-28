package rest

import (
	"context"
	"gogo-gadget-weather/serialize"
	"net/http"
	"time"
)

type Rest[T any] interface {
	Get(client *http.Client, ctx context.Context, url string) (*T, error)
}

func Get[T any](
	client *http.Client,
	ctx context.Context,
	url string,
	timeout time.Duration,
) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)

	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	r, err := serialize.Decode[T](response)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
