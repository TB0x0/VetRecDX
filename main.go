package main

import (
	"fmt"
	"log"
	"net/http"
	"vetrecdx/internal/auth"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", defaultHandler)

	//auth
	mux.HandleFunc("GET /admin", auth.GetUserHandler)

	fmt.Printf("Server listening to %s:%s\n", "localhost", "8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
