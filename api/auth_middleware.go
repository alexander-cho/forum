package main

import (
	"fmt"
	"net/http"

	"github.com/alexander-cho/manager/api/internal/auth"
	"github.com/alexander-cho/manager/api/internal/database"
)

type authenticatedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authorization error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could not retreive user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
