package tests

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"mockroblog/pkg/routes"
)

func TestHello(t *testing.T) {
	testcases := []struct{
		description string
		method string
		path string
		status int
		body string
	}{
		{
			description: "GET request on hello endpoint",
			method: "GET",
			path: "/hello",
			status: 200,
			body: "Hello, World!\n",
		},
		{
			description: "GET request on hello endpoint with trailing /",
			method: "GET",
			path: "/hello/",
			status: 404,
			body: "404 page not found\n",
		},
		{
			description: "POST request on /hello",
			method: "POST",
			path: "/hello",
			status: 405,
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
				t.Errorf("Incorrect response code: %v", recorder.Code)
			}

			if recorder.Body.String() != tc.body {
				t.Errorf("Incorrect response body: %v", recorder.Body.String())
			}

		})
	}
}