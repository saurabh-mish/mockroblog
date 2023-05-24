package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"mockroblog/pkg/routes"
)

func TestHelloHandler(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		status int
		header string
		body string
	}{
		{
			description: "GET request on hello endpoint",
			method: "GET",
			path: "/hello",
			status: 200,
			header: "text/plain",
			body: "Hello, World!\n",
		},
		{
			description: "bad request on hello endpoint - with trailing /",
			method: "GET",
			path: "/hello/",
			status: 404,
			header: "text/plain; charset=utf-8",
			body: "404 page not found\n",
		},
		{
			description: "incorrect method on hello endpoint",
			method: "POST",
			path: "/hello",
			status: 405,
			header: "text/plain; charset=utf-8",
			body: "405 method not allowed\n",
		},
	}

	for _, tc := range testcases {
		t.Logf("testing %s endpoint with %s request", tc.path, tc.method)

		t.Run(tc.description, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatalf("Error receiving response: %v", err)
			}

			routes.Serve(recorder, request)

			if recorder.Code != tc.status {
				t.Errorf("Incorrect code for %v: %v", tc.description, recorder.Code)
			}

			if recorder.Result().Header.Get("Content-Type") != tc.header {
				t.Errorf("Incorrect header for %v: got %v, want %v", tc.description, recorder.Result().Header.Get("Content-Type"), tc.header)
			}

			if recorder.Body.String() != tc.body {
				t.Errorf("Incorrect body for %v: %v", tc.description, recorder.Body.String())
			}

		})
	}
}
