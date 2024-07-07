package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/Mafaz03/rssagg/internal/auth"
	"github.com/Mafaz03/rssagg/internal/database"
)

func (apc *apiConfig) GetUsersByKey(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Key string `json:"api_key"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt Get users by API key: %v", err))
		return
	}
	user, err := apc.DB.GetUserByAPIKey(r.Context(), params.Key)

	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt get the output: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, DatabaseUsertoUser(user))
}

func (apc *apiConfig) GetUsersByAuth(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusCreated, DatabaseUsertoUser(user))
}
