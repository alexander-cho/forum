package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alexander-cho/manager/api/internal/database"
	"github.com/google/uuid"
)

// define http handler to see if server is live and running
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// handler needs to take as input some json body
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	// decode into an instance of the parameter struct, handle error if there is any
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	// otherwise, use db to create user
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	// respond with user object
	respondWithJSON(w, 200, databaseUserToUser(user))
}
