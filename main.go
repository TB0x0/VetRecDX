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
	env_err := godotenv.Load()
	if env_err != nil {
		log.Panicln("Error loading .env file")
	}

	authConfig, err := auth.CreateAuthConfig()
	if err != nil {
		log.Fatal(err)
	}

	db.ConnectDB()
	//db.CreateUser("test2","HASH!!","Assistant") //TEST
	db.GetUser("test2")

	//routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", defaultHandler)
	mux.HandleFunc("GET /health", healthHandler)

	//auth
	mux.HandleFunc("POST /create", auth.ConstructJWT(authConfig))
	mux.HandleFunc("GET /deconstruct/{token}", auth.DeconstructJWT(authConfig)) //for testing, remove in prod

	fmt.Printf("Server listening to %s:%s\n", "localhost", "8080")
	err1 := http.ListenAndServe(":8080", mux)
	if err1 != nil {
		log.Fatal(err1)
	}
}
