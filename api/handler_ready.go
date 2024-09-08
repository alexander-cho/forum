package main

import "net/http"

// define http handler to see if server is live and running
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
