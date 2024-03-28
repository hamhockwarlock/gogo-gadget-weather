package openweathermap

import (
	"net/http"
	"time"
)

type v1 struct {
	baseURL string
	client  http.Client
	timeout time.Duration
}
