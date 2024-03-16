package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/noazaj/Personal-Website/backend/pkg/router"
)

func main() {
	// Get the port to be used for the server
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT number not set")
	}

	// Create a server
	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router.NewRouter(),
	}

	// Start the server on designated port
	fmt.Printf("Starting server on port :%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
