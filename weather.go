package main

import (
	"fmt"
	"net/http"
)

func handleWeatherGet() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("made it to handleWeatherGet!")
		},
	)
}
