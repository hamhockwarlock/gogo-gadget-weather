package main

import (
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
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusUnauthorized,
		)
	}

	parsed, err := decode[ApiErrorResponse](rr.Result())
	if err != nil {
		t.Fatal(err)
	}

	if success := parsed.Success; success != false {
		t.Errorf("handler returned the wrong success: got %t want %t", success, false)
	}

	if errorMessage := parsed.Error; errorMessage != http.StatusText(http.StatusUnauthorized) {
		t.Errorf(
			"handler returned the wrong error: got %s want %s",
			errorMessage,
			http.StatusText(http.StatusUnauthorized),
		)
	}

	if message := parsed.Message; message != apiKeyRequiredMessage {
		t.Errorf(
			"handler returned the wrong message: got %s want %s",
			message,
			apiKeyRequiredMessage,
		)
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
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
