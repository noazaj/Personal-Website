package router

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/noazaj/Personal-Website/backend/pkg/api"
	"github.com/noazaj/Personal-Website/backend/pkg/middleware"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.MiddlewareCors)

	// Serve static files
	publicDir := os.Getenv("PUBLIC_DIR")
	if publicDir == "" {
		log.Fatal("PUBLIC_DIR environment variable is not set")
	}

	router.Handle("/*", middleware.MiddlewarePaths(http.FileServer(http.Dir(publicDir))))

	// Create the subrouter and mount it to the main router
	v1 := chi.NewRouter()
	router.Mount("/v1", v1)

	// Define the routes and associate them
	// with the handlers from the api package
	router.Get("/", api.HandleIndexPage)

	v1.Get("/readiness", api.HandlerReadiness)

	return router
}
