package main

import "net/http"

// define http handler to see if server is live and running
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
