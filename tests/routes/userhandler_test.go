package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"

	"mockroblog/pkg/models"
	"mockroblog/pkg/routes"
)

func TestGetAllUsers(t *testing.T) {
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

			if recorder.Code == 200 {
				var userData models.Users
				err = json.NewDecoder(recorder.Body).Decode(&userData)
				if err != nil {
					t.Errorf("Unable to parse response to structure: got %v, want %v", recorder.Body, userData)
				}
			}
		})
	}
}


func TestCreateUser(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		req_header string
		resp_header string
		body models.User
		status int
	}{
		{
			description: "POST request with correct data",
			method: "POST",
			path: "/api/v1/user",
			req_header: "application/json",
			body: models.User{Username: "username", Password: "p@ssword", Email: "user@domain.com"},
			resp_header: "application/json",
			status: 200,
		},
		{
			description: "POST request with short username",
			method: "POST",
			path: "/api/v1/user",
			req_header: "application/json",
			body: models.User{Username: "user", Password: "p@ssword", Email: "user@domain.com"},
			resp_header: "text/plain; charset=utf-8",
			status: 422,
		},
		{
			description: "POST request with no special character in password",
			method: "POST",
			path: "/api/v1/user",
			req_header: "application/json",
			body: models.User{Username: "user1", Password: "password", Email: "user@domain.com"},
			resp_header: "text/plain; charset=utf-8",
			status: 422,
		},
		{
			description: "POST request with improper email",
			method: "POST",
			path: "/api/v1/user",
			req_header: "application/json",
			body: models.User{Username: "user1", Password: "p@ssword", Email: "user@domain"},
			resp_header: "text/plain; charset=utf-8",
			status: 422,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(tc.body)

			request, err := http.NewRequest(tc.method, tc.path, payloadBuf)
			if err != nil {
				t.Fatalf("Error receiving response: %v", err)
			}
			request.Header.Set("Content-Type", tc.req_header)

			routes.Serve(recorder, request)

			if recorder.Code != tc.status {
				t.Errorf("Incorrect response code for %v: got %v, want %v", tc.description, recorder.Code, tc.status)
			}

			if recorder.Result().Header.Get("Content-Type") != tc.resp_header {
				t.Errorf("Incorrect header for %s: got %v, want %v", tc.description, recorder.Result().Header.Get("Content-Type"), tc.resp_header)
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
			description: "Non-existing user ID",
			method: "GET",
			path: "/api/v1/user/10",
			status: 404,
			header: "text/plain",
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
