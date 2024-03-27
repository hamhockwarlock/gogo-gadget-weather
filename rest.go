package main

import "net/http"

func get[T any](url string) (*T, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
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

	r, err := decode[T](response)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
