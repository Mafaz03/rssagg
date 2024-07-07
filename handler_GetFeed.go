package main

import (
	"fmt"
	"net/http"
)

func (apc *apiConfig) handler_GetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := apc.DB.GetFeeds(r.Context())
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt get the feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, feeds)

}
