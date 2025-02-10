package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/Chance093/chirpy/internal/database"
	"github.com/google/uuid"
)

type apiConfig struct {
	fileServerHits atomic.Int32
	db             *database.Queries
	env            string
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(fmt.Sprintf(`
    <!DOCTYPE html>
		<html>
    <body>
      <h1>Welcome, Chirpy Admin</h1>
      <p>Chirpy has been visited %d times!</p>
    </body>
    </html>
	`, cfg.fileServerHits.Load())))
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
  if cfg.env != "development" {
    respondWithError(w, http.StatusForbidden, "Route forbidden", nil)
    return 
  }

	cfg.fileServerHits.Store(0)

  if err := cfg.db.DeleteAllUsers(r.Context()); err != nil {
    respondWithError(w, http.StatusInternalServerError, "failed to delete users", err)
    return
  }
  
  respondWithJSON(w, http.StatusOK, "Users deleted")
}

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
		respondWithError(w, http.StatusInternalServerError, "error while creating user in db", err)
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
