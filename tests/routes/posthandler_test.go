package routes

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"mockroblog/pkg/models"
// 	"mockroblog/pkg/routes"
// )

// func TestGetAllPosts(t *testing.T) {
// 	recorder := httptest.NewRecorder()
// 	request, err := http.NewRequest("GET", "/api/v1/posts", nil)
// 	if err != nil {
// 		t.Fatalf("Error receiving response: %v", err)
// 	}

// 	routes.Serve(recorder, request)

// 	t.Run("status code", func(t *testing.T){
// 		got := recorder.Code
// 		want := 200
// 		if got != want {
// 			t.Errorf("Incorrect status code received: got %v, want %v", got, want)
// 		}
// 	})

// 	t.Run("header", func(t *testing.T){
// 		got := recorder.Result().Header.Get("Content-Type")
// 		want := "application/json"
// 		if got != want {
// 			t.Errorf("Incorrect header received: got %v, want %v", got, want)
// 		}
// 	})

// 	t.Run("response body", func(t *testing.T){
// 		var allPostsJSON models.Posts
// 		allPostsResponse, _ := ioutil.ReadAll(recorder.Body)

// 		err := json.Unmarshal([]byte(allPostsResponse), &allPostsJSON)
// 		if err != nil {
// 			t.Errorf("Unable to parse response from server: got %v", err)
// 		}
// 	})
// }
