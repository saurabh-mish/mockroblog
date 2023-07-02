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

// func CreateUserWithQueryParams(w http.ResponseWriter, r *http.Request) {
// 	// for param, value := range r.URL.Query() {
// 	// 	log.Println(param, value)
// 	// }

// 	username := r.URL.Query().Get("username")
// 	password := r.URL.Query().Get("password")
// 	email := r.URL.Query().Get("email")

// 	validData := utils.ValidateCreateUser(username, password, email)

// 	if !validData {
// 		http.Error(w, "Could not validate user data", http.StatusUnprocessableEntity)
// 		return
// 	} else {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)

// 		_, err := w.Write([]byte("Welcome " + username + "!\nWe have sent a confirmation email to " + email + "\n"))
// 		if err != nil {
// 			http.Error(w, "Internal server error", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }


func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Request not acceptable; check header", http.StatusNotAcceptable)
	}

	var userData models.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Could not parse user payload", http.StatusUnprocessableEntity)
	}

	_, err = utils.ValidateCreateUser(userData.Username, userData.Password, userData.Email)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnprocessableEntity)
	} else {
		//fmt.Fprintf(w, "User created:\n%v", userData)
		allUsersJSON = append(allUsersJSON, userData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte("Welcome " + userData.Username + "!\nWe have sent a confirmation email to " + userData.Email + "\n"))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}


func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(getField(r, 0), 10, 16)
	// regexp will take integers from 0 to 9; below code unreachable
	// if err != nil {
	// 	http.Error(w, "Unable to convert ID to int", http.StatusUnprocessableEntity)
	// }

	var found bool = false

	for _, user := range allUsersJSON {
		if uint64(user.Id) == userId {
			found = true
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}
	}

	if found == false && userId >= 0 {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
	}
}
