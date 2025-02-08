package main

import (
	"encoding/json"
	"net/http"
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

  // send valid true response
  type returnVals struct {
    Valid bool `json:"valid"`
  }

	resBody := returnVals{
		Valid: true,
	}

  respondWithJSON(w, http.StatusOK, resBody)
}
