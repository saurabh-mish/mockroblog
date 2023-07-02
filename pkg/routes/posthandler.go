package routes

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"mockroblog/pkg/models"
)

var allPosts = models.Posts{
	{14, "title 1", "filler content for post 1", "community 1"},
	{15, "title 2", "filler content for post 2", "community 2"},
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	// r.ParseForm()
	// _, hasNumber := r.Form["number"]
	// _, hasNumberAndCommunity := r.Form["community"]

	// switch {
	// case hasNumberAndCommunity:
	// 	RetrieveMostRecentPostsCommunity(w, r)
	// case hasNumber:
	// 	RetrieveMostRecentPosts(w, r)
	// default:
	// 	http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// switch {
	// case r.URL.Query().Get("number") != "" && r.URL.Query().Get("community") != "":
	// 	RetrieveMostRecentPostsCommunity(w, r)
	// case r.URL.Query().Get("number") != "":
	// 	RetrieveMostRecentPosts(w, r)
	// default:
	// 	http.Error(w, "Invalid query parameter for posts", http.StatusBadRequest)
	// 	return
	// }

	// allPostsJSON, _ := json.Marshal(allPosts)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// _, err := w.Write([]byte(allPostsJSON))
	// if err != nil {
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// }
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Create post\n")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete post\n")
}

func RetrievePost(w http.ResponseWriter, r *http.Request) {
	pId, _ := strconv.Atoi(getField(r, 0))

	var found bool = false

	for _, post := range allPosts {
		if post.Id == pId {
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
