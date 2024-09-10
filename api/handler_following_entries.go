package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alexander-cho/manager/api/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFollowingEntry(w http.ResponseWriter, r *http.Request, user database.User) {
	// handler needs to take as input some json body
	type parameters struct {
		EntryID uuid.UUID `json:"entry_id"`
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
	followingEntry, err := apiCfg.DB.CreateFollowingEntry(r.Context(), database.CreateFollowingEntryParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		EntryID:   params.EntryID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create a following entry: %v", err))
		return
	}

	// respond with user object
	respondWithJSON(w, 201, databaseFollowingEntryToFollowingEntry(followingEntry))
}

// getting all of the entries that a certain user follows
func (apiCfg *apiConfig) handlerGetFollowingEntries(w http.ResponseWriter, r *http.Request, user database.User) {
	// otherwise, use db to create user
	followingEntries, err := apiCfg.DB.GetFollowingEntries(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get following entries: %v", err))
		return
	}

	// respond with user object
	respondWithJSON(w, 201, databaseFollowingEntriesToFollowingEntries(followingEntries))
}

// delete following entry
func (apiCfg *apiConfig) handlerDeleteFollowingEntry(w http.ResponseWriter, r *http.Request, user database.User) {
	followingEntryIDStr := chi.URLParam(r, "followingEntryID")
	followingEntryID, err := uuid.Parse(followingEntryIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not parse feed follow id: %v", err))
	}

	err = apiCfg.DB.DeleteFollowingEntry(r.Context(), database.DeleteFollowingEntryParams{
		ID:     followingEntryID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not delete following feed: %v", err))
	}

	respondWithJSON(w, 200, struct{}{})
}
