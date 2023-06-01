package utils

import (
	"strings"
)

func ValidateCreateUser(username, password, email string) bool {
	if len(username) < 5 || len(username) > 20 {
		return false
	} else if len(password) < 6 || !strings.ContainsAny(password, "@!$#&*") {
		return false
	} else if !strings.ContainsAny(email, "@") || !strings.ContainsAny(email, ".com") {
		return false
	} else {
		return true
	}
}
