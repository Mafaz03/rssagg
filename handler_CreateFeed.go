package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mafaz03/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apc *apiConfig) handler_CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt Decode: %v", err))
		return
	}
	// (id, created_at, updated_at, name, url, user_id)
	feeds, err := apc.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt get the output: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, FeedsUsertoUser(feeds))

}
