package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Chance093/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}

	type response struct {
		User
		Token string `json:"token"`
	}

  // get payload
	pl := payload{}
	if err := json.NewDecoder(r.Body).Decode(&pl); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while decoding payload", err)
		return
	}

  // get user
	user, err := cfg.db.GetUserByEmail(r.Context(), pl.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, http.StatusUnauthorized, "incorrect email or password", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "error while retrieving user", err)
		return
	}

  // check password
	if err = auth.CheckPasswordHash(pl.Password, user.HashedPassword); err != nil {
		respondWithError(w, http.StatusUnauthorized, "incorrect email or password", nil)
		return
	}

  // make jwt token
  expirationTime := time.Hour
	if pl.ExpiresInSeconds > 0 && pl.ExpiresInSeconds < 3600 {
		expirationTime = time.Duration(pl.ExpiresInSeconds) * time.Second
	}

	token, err := auth.MakeJWT(user.ID, cfg.tokenSecret, expirationTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while making jwt", err)
    return
	}

  // respond with user and jwt token
	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token: token,
	})
}
