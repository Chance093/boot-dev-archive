package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Chance093/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) { // Parse payload
  // get payload
	type payload struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	pl := payload{}
	if err := json.NewDecoder(r.Body).Decode(&pl); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while decoding payload", err)
		return
	}

	// validate chirp
  validatedBody, err := validateChirp(pl.Body)
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "", err)
  }

  // create chirp
	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   validatedBody,
		UserID: pl.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while creating chirp", err)
    return
	}

  // send response
	type response struct {
		ID        uuid.UUID `json:"id"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UserID    uuid.UUID `json:"user_id"`
	}

	respondWithJSON(w, http.StatusCreated, response{
    ID: chirp.ID,
    Body: chirp.Body,
    CreatedAt: chirp.CreatedAt,
    UpdatedAt: chirp.UpdatedAt,
    UserID: chirp.UserID,
	})
}

func validateChirp(body string) (string, error) {
	if len(body) > 140 {
		return "", errors.New("chirp is too long")
	}

	return removeProfan(body), nil
}

func removeProfan(body string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(body,
							"Fornax", "****"),
						"Sharbert", "****"),
					"Kerfuffle", "****"),
				"fornax", "****"),
			"sharbert", "****"),
		"kerfuffle", "****")
}
