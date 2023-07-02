package routes

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"mockroblog/pkg/models"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	allPosts := models.Posts{
		{nil, "title 1", "filler content for post 1", "community 1"},
		{nil, "title 2", "filler content for post 2", "community 2"},
	}

	allPostsJSON, _ := json.Marshal(allPosts)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(allPostsJSON))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
