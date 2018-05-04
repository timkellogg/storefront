package main

import (
	"encoding/json"
	"net/http"
)

// respondWithError - send back http response with error to web client
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// respondWithJSON - send back JSON response to web client
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
