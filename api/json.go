package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respond with arbitrary error messages
func respondWithError(w http.ResponseWriter, code int, message string) {
	// server side errors
	if code > 499 {
		log.Println("Responding with 5xx error:", message)
	}

	// structure the JSON response sent back to the client when error occurs
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: message,
	})
}

// serialize data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("failed to marshal json response %v", payload)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
