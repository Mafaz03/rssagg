package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mafaz03/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apc *apiConfig) handler_CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		feed_id uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt Decode: %v", err))
		return
	}
	// (id, created_at, updated_at, name, url, user_id)
	feeds, err := apc.DB.CreateFeedsFollow(r.Context(), database.CreateFeedsFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.feed_id,
	})
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt get the output: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, feeds)

}
