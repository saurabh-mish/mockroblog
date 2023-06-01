package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

	"mockroblog/pkg/models"
	"mockroblog/pkg/utils"
)


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
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
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// for param, value := range r.URL.Query() {
	// 	log.Println(param, value)
	// }

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	email := r.URL.Query().Get("email")

	validData := utils.ValidateCreateUser(username, password, email)

	if !validData {
		http.Error(w, "Could not validate user data", http.StatusUnprocessableEntity)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("Welcome " + username + "!\nWe have sent a confirmation email to " + email + "\n"))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(getField(r, 0), 10, 16)
	if err != nil {
		http.Error(w, "Unable to convert ID to int", http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "The value of ID is %v\n", id)
}
