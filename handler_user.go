package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Chance093/chirpy/internal/auth"
	"github.com/Chance093/chirpy/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
  // get payload
	type payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	pl := payload{}
	if err := json.NewDecoder(r.Body).Decode(&pl); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while decoding payload", err)
		return
	}

  // validate email
	if pl.Email == "" {
		respondWithError(w, http.StatusBadRequest, "payload missing email", nil)
		return
	}

  // hash password
	hashedPassword, err := auth.HashPassword(pl.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while hashing password", err)
	}

  // create user in db
	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Email:          pl.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while creating user", err)
		return
	}

  // respond with created user
	respondWithJSON(w, http.StatusCreated, User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
