package middleware

import (
	"net/http"
	"strings"

	"github.com/noazaj/Personal-Website/backend/pkg/util"
)

func MiddlewarePaths(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if strings.HasPrefix(path, "/images") || strings.HasPrefix(path, "/css") {
			if isDirectFilePath(path) {
				next.ServeHTTP(w, r)
			} else {
				util.RespondWithError(w, http.StatusForbidden, "Forbidden Directory")
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func isDirectFilePath(path string) bool {
	if strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".css") {
		return true
	}
	return false
}
