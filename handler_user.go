package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Email string `json:"email"`
	}

	pl := payload{}
	if err := json.NewDecoder(r.Body).Decode(&pl); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while decoding payload", err)
		return
	}

	if pl.Email == "" {
		respondWithError(w, http.StatusBadRequest, "payload missing email", nil)
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), pl.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while creating user", err)
		return
	}

	type response struct {
		ID        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	respondWithJSON(w, http.StatusCreated, response{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
