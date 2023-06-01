package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"mockroblog/pkg/models"
	"mockroblog/pkg/routes"
)

func TestGetAllUsersForHeader(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		status int
		header string
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


func TestGetAllUsersForData(t *testing.T) {
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


func TestCreateUser(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		status int
		header string
	}{
		{
			description: "POST request with correct user data",
			method: "POST",
			path: "/api/v1/user?username=testuser&email=testuser@domain.com&password=p@ssword",
			status: 200,
			header: "application/json",
		},
		{
			description: "POST request with wrong username format",
			method: "POST",
			path: "/api/v1/user?username=user&email=testuser@domain.com&password=p@ssword",
			status: 422,
			header: "text/plain; charset=utf-8",
		},
		{
			description: "POST request with no special character in password",
			method: "POST",
			path: "/api/v1/user?username=testuser&email=testuser@domain.com&password=password",
			status: 422,
			header: "text/plain; charset=utf-8",
		},
		{
			description: "POST request with improper email",
			method: "POST",
			path: "/api/v1/user?username=testuser&email=testuserdomain&password=p@ssword",
			status: 422,
			header: "text/plain; charset=utf-8",
		},
		{
			description: "GET request with correct user data",
			method: "GET",
			path: "/api/v1/user?username=testuser&email=testuser@domain.com&password=p@ssword",
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


func TestRetrieveUser(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		status int
		header string
	}{
		{
			description: "Valid user ID",
			method: "GET",
			path: "/api/v1/user/12",
			status: 200,
			header: "application/json",
		},
		{
			description: "Invalid user ID - negative number",
			method: "GET",
			path: "/api/v1/user/-35",
			status: 404,
			header: "text/plain; charset=utf-8",
		},
		{
			description: "Invalid user ID - out of range",
			method: "GET",
			path: "/api/v1/user/450293010",
			status: 422,
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
				t.Errorf("Incorrect response code: got %v, want %v", recorder.Code, tc.status)
			}

			if recorder.Result().Header.Get("Content-Type") != tc.header {
				t.Errorf("Incorrect header for %s: got %v, want %v", tc.description, recorder.Result().Header.Get("Content-Type"), tc.header)
			}
		})
	}
}
