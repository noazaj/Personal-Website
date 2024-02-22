package api

import (
	"net/http"
	"os"
	"path/filepath"
)

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	publicDir := filepath.Join(workDir, "../../../frontend/public/index.html")
	http.ServeFile(w, r, publicDir)
}
