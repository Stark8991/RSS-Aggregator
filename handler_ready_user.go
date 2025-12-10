package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Stark8991/RSSAgg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json: "name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: ", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func handlerCreateUserErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (cfg *apiConfig) handlUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))

}
