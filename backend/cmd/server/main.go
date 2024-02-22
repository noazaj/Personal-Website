package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error: couldn't load .env variables")
	}

	// Get the port to be used for the server
	port := os.Getenv("PORT")

	// Create a server
	srv := &http.Server{
		Addr: port,
	}

	// Start the server on designated port
	fmt.Printf("Starting server on port :%s", port)
	log.Fatal(srv.ListenAndServe())
}
