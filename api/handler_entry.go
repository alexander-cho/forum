package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alexander-cho/manager/api/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateEntry(w http.ResponseWriter, r *http.Request, user database.User) {
	// handler needs to take as input some json body
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
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
	entry, err := apiCfg.DB.CreateEntry(r.Context(), database.CreateEntryParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create entry: %v", err))
		return
	}

	// respond with user object
	respondWithJSON(w, 201, databaseEntryToEntry(entry))
}

// no auth necessary
func (apiCfg *apiConfig) handlerGetEntries(w http.ResponseWriter, r *http.Request) {

	entries, err := apiCfg.DB.GetEntries(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not retrieve entries: %v", err))
		return
	}

	// respond with user object
	respondWithJSON(w, 201, databaseEntriesToEntries(entries))
}
