package main

import (
	"fmt"
	"net/http"

	"github.com/Mafaz03/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apc *apiConfig) DeleteFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedFollowIdStr := chi.URLParam(r, "feedfollowid")
	FeedFollowId, err := uuid.Parse(FeedFollowIdStr)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt get the Delete statement: %v", err))
	}
	err = apc.DB.DeleteFeedsFollow(r.Context(), FeedFollowId)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt Delete: %v", err))
	}

	respondWithJSON(w, http.StatusCreated, struct{}{})
}
