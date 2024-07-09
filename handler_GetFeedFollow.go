package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (apc *apiConfig) handler_GetFeedFollow(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		User_id uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt Decode: %v", err))
		return
	}
	feedsfollow, err := apc.DB.GetFeedsFollow(r.Context(), params.User_id)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt get the output: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, feedsfollow)

}
