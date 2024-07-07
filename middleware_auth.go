package main

import (
	"fmt"
	"net/http"

	"github.com/Mafaz03/rssagg/internal/auth"
	"github.com/Mafaz03/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetApiKey(r.Header)
		if err != nil {
			errResponse(w, 403, fmt.Sprintf("Auth err: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), api_key)
		if err != nil {
			errResponse(w, 400, fmt.Sprintf("Couldnt Query for api key: %v", err))
			return
		}
		handler(w, r, user)
	}
}

