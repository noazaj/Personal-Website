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
			if isDirectFileAccess(path) {
				util.RespondWithError(w, http.StatusForbidden, "Forbidden")
			} else {
				next.ServeHTTP(w, r)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func isDirectFileAccess(path string) bool {
	splitPath := strings.Split(path, "/")
	final := splitPath[len(splitPath)-1]
	isFile := false
	for _, char := range final {
		if char == '.' {
			isFile = true
			break
		}
	}
	return isFile
}
