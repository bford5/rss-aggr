package main

import (
	"fmt"
	"net/http"

	"github.com/bford5/rss-aggr/internal/auth"
	"github.com/bford5/rss-aggr/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("failed to get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
