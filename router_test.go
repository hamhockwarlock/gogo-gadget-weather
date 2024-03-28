package main

import (
	"fmt"
	"gogo-gadget-weather/customerror"
	"gogo-gadget-weather/serialize"
	"gogo-gadget-weather/test"
	"gogo-gadget-weather/weather"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleNotFound(t *testing.T) {
	uri := "/api/nope"
	req := httptest.NewRequest("GET", uri, nil)
	w := weather.New(nil)

	req.Header.Add(apiKeyHeaderName, "anything")
	rr := httptest.NewRecorder()
	r := router(w)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		test.Fatal(t, status, http.StatusNotFound)
	}

	parsed, err := serialize.Decode[customerror.ApiErrorResponse](rr.Result())
	if err != nil {
		t.Fatal(err)
	}

	if success := parsed.Success; success != false {
		test.Fatal(t, success, false)
	}

	if errorMessage := parsed.Error; errorMessage != http.StatusText(http.StatusNotFound) {
		test.Fatal(t, errorMessage, http.StatusText(http.StatusNotFound))
	}

	expectedMessage := fmt.Sprintf("No endpoint at %s", uri)
	if message := parsed.Message; message != expectedMessage {
		test.Fatal(t, message, expectedMessage)
	}
}

func TestHandleMethodNotAllowed(t *testing.T) {
	uri := "/api/weather"
	method := "DELETE"
	req := httptest.NewRequest(method, uri, nil)
	w := weather.New(nil)

	req.Header.Add(apiKeyHeaderName, "anything")
	rr := httptest.NewRecorder()
	r := router(w)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		test.Fatal(t, status, http.StatusMethodNotAllowed)
	}

	parsed, err := serialize.Decode[customerror.ApiErrorResponse](rr.Result())
	if err != nil {
		t.Fatal(err)
	}

	if success := parsed.Success; success != false {
		test.Fatal(t, success, false)
	}

	if errorMessage := parsed.Error; errorMessage != http.StatusText(http.StatusMethodNotAllowed) {
		test.Fatal(t, errorMessage, http.StatusText(http.StatusNotFound))
	}

	expectedMessage := fmt.Sprintf("%s not allowed for endpoint", method)

	if message := parsed.Message; message != expectedMessage {
		test.Fatal(t, message, expectedMessage)
	}
}
