package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Chance093/chirpy/internal/auth"
	"github.com/Chance093/chirpy/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	// validate jwt
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid authorization header", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	// get payload
	type payload struct {
		Body string `json:"body"`
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
		return
	}

	// create chirp
	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   validatedBody,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while creating chirp", err)
		return
	}

	// send response
	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		Body:      chirp.Body,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserID:    chirp.UserID,
	})
}

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	// extract optional authorId
	queryParams := r.URL.Query()
	authorIDStr := queryParams.Get("author_id")
	var authorId uuid.UUID
	if authorIDStr != "" {
		authorID, err := uuid.Parse(authorIDStr)
		if err != nil {
			http.Error(w, "Invalid author_id", http.StatusBadRequest)
			return
		}
		authorId = authorID
	}

	var chirps []database.Chirp
	var err error
  if len(authorId) > 0 {
		chirps, err = cfg.db.GetAllChirpsByAuthorId(r.Context(), authorId)
	} else {
		chirps, err = cfg.db.GetAllChrips(r.Context())
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while getting all chirps", err)
	}

	res := make([]Chirp, 0, len(chirps))

	for i := 0; i < len(chirps); i++ {
		res = append(res, Chirp{
			ID:        chirps[i].ID,
			Body:      chirps[i].Body,
			CreatedAt: chirps[i].CreatedAt,
			UpdatedAt: chirps[i].UpdatedAt,
			UserID:    chirps[i].UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, res)
}

func (cfg *apiConfig) handlerGetChirpByID(w http.ResponseWriter, r *http.Request) {
	chirpIdString := r.PathValue("id")

	// parse id into UUID
	chirpID, err := uuid.Parse(chirpIdString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp id", err)
		return
	}

	// search database for chirp
	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, http.StatusNotFound, "no chirp with that id exists", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "error while retrieving chirp", err)
		return
	}

	// send chirp as response
	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		Body:      chirp.Body,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserID:    chirp.UserID,
	})
}

func (cfg *apiConfig) handlerDeleteChirpByID(w http.ResponseWriter, r *http.Request) {
	// validate jwt
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid authorization header", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	chirpIdString := r.PathValue("id")

	// parse id into UUID
	chirpID, err := uuid.Parse(chirpIdString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp id", err)
		return
	}

	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, http.StatusNotFound, "no chirp with that id exists", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "error while retrieving chirp", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "forbidden resource", err)
		return
	}

	if err := cfg.db.DeleteChirpByID(r.Context(), chirpID); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while deleting chirp", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
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
