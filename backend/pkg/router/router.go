package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/zajicekn/Personal-Website/pkg/api"
	"github.com/zajicekn/Personal-Website/pkg/middleware"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.MiddlewareCors)

	// Serve static files
	workDir, _ := os.Getwd()
	publicDir := filepath.Join(workDir, "../../frontend/public")
	router.Handle("/*", http.FileServer(http.Dir(publicDir)))

	// Create the subrouter and mount it to the main router
	v1 := chi.NewRouter()
	router.Mount("/v1", v1)

	// Define the routes and associate them
	// with the handlers from the api package
	v1.Get("/readiness", api.HandlerReadiness)
	v1.Get("/", api.HandleIndexPage)

	return router
}
