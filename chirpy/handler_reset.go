package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.env != "development" {
		respondWithError(w, http.StatusForbidden, "reset is only allowed in dev environment.", nil)
		return
	}

	cfg.fileServerHits.Store(0)

	if err := cfg.db.DeleteAllUsers(r.Context()); err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to delete users", err)
		return
	}

	respondWithJSON(w, http.StatusOK, "hits set to 0 and users deleted")
}
