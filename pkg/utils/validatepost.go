package utils


import (
	"errors"
	//"strings"
)

func ValidatePostData(title, text, community string) (bool, error) {
	if len(title) < 5 || len(text) < 20 || len(community) < 3 {
		return false, errors.New("Length of title, text, and / or community should be higher")
	} else {
		return true, nil
	}
}
