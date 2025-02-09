package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Chance093/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
  const PORT = "8080"
  
  // load .env
  if err := godotenv.Load(); err != nil {
    log.Fatalf("could not load .env: %v", err)
  }


  // connect db
  dbUrl := os.Getenv("DB_URL")
  if dbUrl == "" {
    log.Fatal("DB_URL must be set in .env")
  }

  db, err := sql.Open("postgres", dbUrl)
  if err != nil {
    log.Fatalf("could not connect to db: %v", err)
  }
  
  // init config
  cfg := &apiConfig{
    fileServerHits: atomic.Int32{},
    db: database.New(db),
  }

  // handler routes
	mux := http.NewServeMux()
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", handlerHealth)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)

  // init server
	s := &http.Server{
		Handler: mux,
		Addr:    ":" + PORT,
	}

  // listen and serve
	fmt.Printf("serving on port :%v\n", PORT)
	log.Fatal(s.ListenAndServe())
}
