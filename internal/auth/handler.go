package auth

import (
	"fmt"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := getUser("admin")

	_, err := fmt.Fprintf(w, "admin has userId: %d", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
