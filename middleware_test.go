package main

import (
	"gogo-gadget-weather/customerror"
	"gogo-gadget-weather/serialize"
	"gogo-gadget-weather/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestApiKeyNotPassed calls apiKeyRequired with a request that did not pass an API key.
// If it is not passed we return a 401.
func TestApiKeyNotPassed(t *testing.T) {

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrappedHandler := apiKeyRequired(dummyHandler)

	req := httptest.NewRequest("GET", "/", nil)

	rr := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		test.Fatal(t, status, http.StatusUnauthorized)
	}

	parsed, err := serialize.Decode[customerror.ApiErrorResponse](rr.Result())
	if err != nil {
		t.Fatal(err)
	}

	if success := parsed.Success; success != false {
		test.Fatal(t, success, false)
	}

	if errorMessage := parsed.Error; errorMessage != http.StatusText(http.StatusUnauthorized) {
		test.Fatal(t, errorMessage, http.StatusText(http.StatusUnauthorized))
	}

	if message := parsed.Message; message != apiKeyRequiredMessage {
		test.Fatal(t, message, apiKeyRequiredMessage)
	}
}

// TestApiKeyPassed calls apiKeyRequired with a request that did pass an API key.
// If it is passed we return a 200.
func TestApiKeyPassed(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrappedHandler := apiKeyRequired(dummyHandler)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add(apiKeyHeaderName, "anything")

	rr := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		test.Fatal(t, status, http.StatusOK)
	}
}
