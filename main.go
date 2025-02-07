package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	const port = "8080"
	cfg := &apiConfig{
		fileServerHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", handlerHealth)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
  mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)

	s := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	fmt.Printf("serving on port :%v\n", port)
	log.Fatal(s.ListenAndServe())
}
