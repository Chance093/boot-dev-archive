package main

import "net/http"

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(http.StatusText(status)))
}
