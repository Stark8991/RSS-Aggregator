package main

import (
	"net/http"

	"github.com/Stark8991/RSSAgg/internal/auth"
	"github.com/Stark8991/RSSAgg/internal/database"
)

type autheHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler autheHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApIKey(r.Header)

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, http.StatusNotFound, "couldn't get user")
			return
		}

		handler(w, r, user)

	}
}
