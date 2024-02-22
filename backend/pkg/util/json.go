package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: could not marshal payload: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errorResponse{Error: msg})
}
