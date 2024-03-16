package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/noazaj/Personal-Website/backend/pkg/util"
)

func MiddlewarePaths(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Println(path)

		if strings.HasPrefix(path, "/images") || strings.HasPrefix(path, "/css") {
			util.RespondWithError(w, http.StatusForbidden, "Forbidden Directory")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
