package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mockroblog/pkg/models"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World!\n")
}


func AllUsers(w http.ResponseWriter, r *http.Request) {
	allUsersJSON := []models.User{
		{12, "user1", "user1@domain.com", "", 2},
		{23, "user2", "user2@domain.com", "", 20},
		{11, "user3", "user3@domain.com", "", 8},
	}

	w.Header().Set("Content-Type", "application/json")

	// http.ResponseWriter.Write implicitly calls w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&allUsersJSON)

	if err != nil {
		http.Error(w, "Unable to parse data", http.StatusInternalServerError)
	}
	//w.WriteHeader(http.StatusOK)    // superfluous response.WriteHeader call
}
