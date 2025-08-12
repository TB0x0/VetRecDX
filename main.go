package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"vetrecdx/internal/auth"
	"vetrecdx/internal/db"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", defaultHandler)
	// DB Connection
	env_err := godotenv.Load()
	if env_err != nil {
		log.Panicln("Error loading .env file")
	}
	db.ConnectDB()

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
