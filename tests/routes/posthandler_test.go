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

func TestCreatePost(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		reqHeader string
		body models.Post
		respHeader string
		status int
	}{
		{
			description: "Create blogpost with correct data",
			method: "POST",
			path: "/api/v1/post",
			reqHeader: "application/json",
			body: models.Post{Title: "Lorem Ipsum", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", Community: "playground"},
			respHeader: "application/json",
			status: 200,
		},
		{
			description: "Create blogpost with correct data and incorrect header",
			method: "POST",
			path: "/api/v1/post",
			reqHeader: "text/plain",
			body: models.Post{Title: "English Pangram", Content: "The quick brown fox jumps over the lazy dog", Community: "playground"},
			respHeader: "text/plain; charset=utf-8",
			status: 406,
		},
		{
			description: "Create blogpost with incorrect payload",
			method: "POST",
			path: "/api/v1/post",
			reqHeader: "application/json",
			body: models.Post{Title: "", Content: "", Community: ""},
			respHeader: "text/plain; charset=utf-8",
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

			request.Header.Set("Content-Type", tc.reqHeader)
			routes.Serve(recorder, request)

			if recorder.Code != tc.status {
				t.Errorf("Incorrect response code for %v: got %v, want %v", tc.description, recorder.Code, tc.status)
			}

			if recorder.Result().Header.Get("Content-Type") != tc.respHeader {
				t.Errorf("Incorrect header for %s: got %v, want %v", tc.description, recorder.Result().Header.Get("Content-Type"), tc.respHeader)
			}
		})
	}
}


func TestRetrievePost(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		header string
		status int
	}{
		{
			description: "Retrieve existing post",
			method: "GET",
			path: "/api/v1/post/id=14",
			header: "application/json",
			status: 200,
		},
		{
			description: "Retrieve non-existing post",
			method: "GET",
			path: "/api/v1/post/id=100",
			header: "text/plain",
			status: 404,
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

				if recorder.Code == 200 {
					var postData models.Posts
					err = json.NewDecoder(recorder.Body).Decode(&postData)
					if err != nil {
						t.Errorf("Unable to parse response to structure: got %v, want %v", recorder.Body, postData)
					}
				}
			}

			if recorder.Result().Header.Get("Content-Type") != tc.header {
				t.Errorf("Incorrect header for %s: got %v, want %v", tc.description, recorder.Result().Header.Get("Content-Type"), tc.header)
			}
		})
	}
}






