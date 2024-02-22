package api

import (
	"net/http"

	"github.com/zajicekn/Personal-Website/pkg/util"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type payload struct {
		Status string `json:"status"`
	}
	util.RespondWithJson(w, http.StatusOK, payload{"ok"})
}
