package api_test

import (
	"BakingApp/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

//First Hello world test
func TestGetRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	request_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.HandleGet)

	handler.ServeHTTP(request_recorder, req)

	if status := request_recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message": "Hello, world!"}`
	if request_recorder.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", request_recorder.Body.String(), expected)
	}
}

//
