package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	dat, err := json.Marshal(payload)
	if err != nil {
    respondWithError(w, http.StatusInternalServerError, "error while marshalling response body", err)
		return
	}

	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	type errResponse struct {
		Error string `json:"error"`
	}

	if code > 499 {
		log.Printf("%s: %s", msg, err)
		w.WriteHeader(code)
		return
	}

	errRes := errResponse{
		Error: msg,
	}

  respondWithJSON(w, code, errRes)
}
