package api

import (
	"net/http"
	"os"
	"path/filepath"
)

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	publicDir := os.Getenv("PUBLIC_DIR")
	indexFilePath := filepath.Join(publicDir, "index.html")
	http.ServeFile(w, r, indexFilePath)
}
