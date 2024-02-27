package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/noazaj/Personal-Website/backend/pkg/router"
)

func main() {
	// Load env variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Couldn't load .env variables. Error: %s", err)
	}

	// Get the port to be used for the server
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT number must be set in .env file")
	}

	// Create a server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router.NewRouter(),
	}

	// Start the server on designated port
	fmt.Printf("Starting server on port :%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
