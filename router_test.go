package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleNotFound(t *testing.T) {
	uri := "/api/nope"
	req := httptest.NewRequest("GET", uri, nil)

	req.Header.Add(apiKeyHeaderName, "anything")
	rr := httptest.NewRecorder()
	r := router()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf(
			"handler returned the wrong status code: got %v want %v",
			status,
			http.StatusNotFound,
		)
	}

	parsed, err := decode[ApiErrorResponse](rr.Result())
	if err != nil {
		t.Fatal(err)
	}

	if success := parsed.Success; success != false {
		t.Errorf("handler returned the wrong success: got %t want %t", success, false)
	}

	if errorMessage := parsed.Error; errorMessage != http.StatusText(http.StatusNotFound) {
		t.Errorf(
			"handler returned the wrong error: got %s want %s",
			errorMessage,
			http.StatusText(http.StatusNotFound),
		)
	}

	expectedMessage := fmt.Sprintf("No endpoint at %s", uri)
	if message := parsed.Message; message != expectedMessage {
		t.Errorf("handler returned the wrong message: got %s want %s", message, expectedMessage)
	}
}

func TestHandleMethodNotAllowed(t *testing.T) {
	uri := "/api/weather"
	method := "DELETE"
	req := httptest.NewRequest(method, uri, nil)

	req.Header.Add(apiKeyHeaderName, "anything")
	rr := httptest.NewRecorder()
	r := router()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf(
			"handler returned the wrong status code: got %v want %v",
			status,
			http.StatusMethodNotAllowed,
		)
	}

	parsed, err := decode[ApiErrorResponse](rr.Result())
	if err != nil {
		t.Fatal(err)
	}

	if success := parsed.Success; success != false {
		t.Errorf("handler returned the wrong success: got %t want %t", success, false)
	}

	if errorMessage := parsed.Error; errorMessage != http.StatusText(http.StatusMethodNotAllowed) {
		t.Errorf(
			"handler returned the wrong error: got %s want %s",
			errorMessage,
			http.StatusText(http.StatusNotFound),
		)
	}

	expectedMessage := fmt.Sprintf("%s not allowed for endpoint", method)

	if message := parsed.Message; message != expectedMessage {
		t.Errorf("handler returned the wrong message: got %s want %s", message, expectedMessage)
	}
}
