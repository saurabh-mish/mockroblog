package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"mockroblog/pkg/models"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello, World!\n"))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func AllUsers(w http.ResponseWriter, r *http.Request) {
	allUsersJSON := []models.User{
		{12, "user1", "user1@domain.com", "", 2},
		{23, "user2", "user2@domain.com", "", 20},
		{11, "user3", "user3@domain.com", "", 8},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&allUsersJSON)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func AllPosts(w http.ResponseWriter, r *http.Request) {
	allPosts := models.Posts{
		{nil, "title 1", "filler content for post 1", "community 1", "user1", "bit.ly/sa12", time.Time{}, 10, 2},
		{nil, "title 2", "filler content for post 2", "community 2", "user1", "bit.ly/sa98", time.Time{}, 11, 4},
	}

	allPostsJSON, _ := json.Marshal(allPosts)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(allPostsJSON))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
