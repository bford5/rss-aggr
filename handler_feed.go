package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bford5/rss-aggr/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json"))
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, dbFeedToFeed(feed))
}

func (apiConfig *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, dbFeedsToFeeds(feeds))
}
