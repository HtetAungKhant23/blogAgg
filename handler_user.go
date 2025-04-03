package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/HtetAungKhant23/blogAgg/internal/auth"
	"github.com/HtetAungKhant23/blogAgg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	param := parameters{}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&param)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      param.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, 200, fmt.Sprintf("Could not get user: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))
}
