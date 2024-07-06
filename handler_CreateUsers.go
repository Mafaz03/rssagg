package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mafaz03/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apc *apiConfig) handler_CreateUsers(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt Decode: %v", err))
		return
	}

	user, err := apc.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Couldnt get the output: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, DatabaseUsertoUser(user))

}
