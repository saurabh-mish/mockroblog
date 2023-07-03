package routes

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"mockroblog/pkg/models"
	"mockroblog/pkg/utils"
)

var allPosts = models.Posts{
	{14, "title 1", "filler content for post 1", "community 1"},
	{15, "title 2", "filler content for post 2", "community 2"},
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Request not acceptable; check header", http.StatusNotAcceptable)
		return
	}

	var postData models.Post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		http.Error(w, "Could not parse post payload", http.StatusUnprocessableEntity)
		return
	}

	_, err = utils.ValidatePostData(postData.Title, postData.Content, postData.Community)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnprocessableEntity)
		return
	} else {
		allPosts = append(allPosts, postData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, "Post created successfully!\n")
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete post\n")
}

func RetrievePost(w http.ResponseWriter, r *http.Request) {
	postId, _ := strconv.Atoi(getField(r, 0))

	var found bool = false

	for _, post := range allPosts {
		if post.Id == postId {
			found = true
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(post)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}
	}

	if found == false {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
	}
}

func RetrieveRecentPosts(w http.ResponseWriter, r *http.Request) {
	num := getField(r, 0)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v Most recent posts!\n", num)
}

func RetrieveRecentPostsFromCommunity(w http.ResponseWriter, r *http.Request) {
	number := getField(r, 0)
	commmunity := getField(r, 1)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v Most recent posts from community %v\n", number, commmunity)
}
