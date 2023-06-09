package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

	"mockroblog/pkg/models"
	"mockroblog/pkg/utils"
)

var allUsersJSON = []models.User{
	{12, "user1", "user1@domain.com", "", 2},
	{23, "user2", "user2@domain.com", "", 20},
	{11, "user3", "user3@domain.com", "", 8},
}


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&allUsersJSON)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Request not acceptable; check header", http.StatusNotAcceptable)
		return
	}

	var userData models.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Could not parse user payload", http.StatusUnprocessableEntity)
		return
	}

	_, err = utils.ValidateUserData(userData.Username, userData.Password, userData.Email)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnprocessableEntity)
		return
	} else {
		allUsersJSON = append(allUsersJSON, userData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "Welcome %v!\nWe have sent a confirmation email to %v", userData.Username, userData.Email)
	}
}


func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(getField(r, 0), 10, 16)
	// regexp takes integers from 0 to 9; _ (err) is unreachable

	var found bool = false

	for _, user := range allUsersJSON {
		if uint64(user.Id) == userId {
			found = true
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				http.Error(w, "Unable to retrieve user", http.StatusInternalServerError)
			}
		}
	}

	if found == false {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
	}
}
