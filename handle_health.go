package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(http.StatusText(status)))
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	// Parse payload
	type payload struct {
		Body string `json:"body"`
	}

	pl := payload{}
	if err := json.NewDecoder(r.Body).Decode(&pl); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while decoding payload", err)
		return
	}

	// check if chirp is too long
	if len(pl.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanedBod := removeProfan(pl.Body)

	respondWithJSON(w, http.StatusOK, struct {
		CleanedBody string `json:"cleaned_body"`
	}{
		CleanedBody: cleanedBod,
	})
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
