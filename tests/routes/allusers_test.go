package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"mockroblog/pkg/models"
	"mockroblog/pkg/routes"
)

func TestAllUsersHandlerForHeader(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		status int
		header string
		body string
	}{
		{
			description: "GET request to all users endpoint",
			method: "GET",
			path: "/api/v1/users",
			status: 200,
			header: "application/json",
		},
		{
			description: "incorrect method on all users endpoint",
			method: "POST",
			path: "/api/v1/users",
			status: 405,
			header: "text/plain; charset=utf-8",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatalf("Error receiving response: %v", err)
			}

			routes.Serve(recorder, request)

			if recorder.Code != tc.status {
				t.Errorf("Incorrect response code for %v: got %v, want %v", tc.description, recorder.Code, tc.status)
			}

			if recorder.Result().Header.Get("Content-Type") != tc.header {
				t.Errorf("Incorrect header for %s: got %v, want %v", tc.description, recorder.Result().Header.Get("Content-Type"), tc.header)
			}
		})
	}
}


func TestAllUsersHandlerForValidData(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Fatalf("Error receiving response: %v", err)
	}

	routes.Serve(recorder, request)

	var got models.Users

	err = json.NewDecoder(recorder.Body).Decode(&got)
	if err != nil {
		t.Errorf("Unable to parse response from server: got %v, want %v", recorder.Body, err)
	}
}
