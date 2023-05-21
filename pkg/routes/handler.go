package routes

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")
}
