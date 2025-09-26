package main

import (
	"encoding/json"
	"net/http"

	"github.com/Chance093/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPolkaWebhook(w http.ResponseWriter, r *http.Request) {
  // validate api key 
  apiKey, err := auth.GetAPIKey(r.Header)
  if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid authorization header", err)
		return
  }
  if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Invalid api key", err)
		return
  }

	// get payload
	type payload struct {
		Event string `json:"event"`
		Data  struct {
			UserId uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	pl := payload{}
	if err := json.NewDecoder(r.Body).Decode(&pl); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while decoding payload", err)
		return
	}

	// don't care about any other event
	if pl.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}

  // update chirpy red membership
	res, err := cfg.db.UpdateChirpyRedActive(r.Context(), pl.Data.UserId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while updating chirpy red membership", err)
		return
	}

  // check how many rows affected
	rows, err := res.RowsAffected()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while getting rows affected", err)
		return
	}
	if rows == 0 {
		respondWithError(w, http.StatusNotFound, "no user with that id exists", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
