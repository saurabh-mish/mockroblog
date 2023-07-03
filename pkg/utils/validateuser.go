package utils


import (
	"strings"
	"errors"
)

func ValidateUserData(username, password, email string) (bool, error) {
	if len(username) < 5 || len(username) > 20 {
		return false, errors.New("Length of username incorrect")
	} else if len(password) < 6 || !strings.ContainsAny(password, "@!$#&*") {
		return false, errors.New("Password too short or doesn't contain special characters")
	} else if !strings.Contains(email, ".com") || !strings.Contains(email, "@") {
		return false, errors.New("Invalid email address")
	} else {
		return true, nil
	}
}
