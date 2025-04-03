package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/HtetAungKhant23/blogAgg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	param := parameters{}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&param)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      param.Name,
		Url:       param.Url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error creating feed: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error getting feeds: %v", err))
	}

	responseWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
