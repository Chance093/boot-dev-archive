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

type apiConfig struct {
	fileServerHits atomic.Int32
	db             *database.Queries
	env            string
	tokenSecret    string
}

func main() {
	const PORT = "8080"

	// load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("could not load .env: %v", err)
	}

	// validate .env vars
	dbUrl := os.Getenv("DB_URL")
	serverEnv := os.Getenv("SERVER_ENV")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	if dbUrl == "" {
		log.Fatal("DB_URL must be set in .env")
	}
	if serverEnv == "" {
		log.Fatal("SERVER_ENV must be set in .env")
	}
	if tokenSecret == "" {
		log.Fatal("TOKEN_SECRET must be set in .env")
	}

	// connect db
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	// init config
	cfg := &apiConfig{
		fileServerHits: atomic.Int32{},
		db:             database.New(dbConn),
		env:            serverEnv,
		tokenSecret:    tokenSecret,
	}

	// handler routes
	mux := http.NewServeMux()
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))

	mux.HandleFunc("GET /api/healthz", handlerHealth)
	mux.HandleFunc("POST /api/login", cfg.handlerLogin)
	mux.HandleFunc("POST /api/refresh", cfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", cfg.handlerRevoke)

	mux.HandleFunc("GET /api/chirps", cfg.handlerGetAllChirps)
	mux.HandleFunc("POST /api/chirps", cfg.handlerCreateChirp)
	mux.HandleFunc("GET /api/chirps/{id}", cfg.handlerGetChirpByID)

	mux.HandleFunc("POST /api/users", cfg.handlerCreateUser)

	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)

	// init server
	s := &http.Server{
		Handler: mux,
		Addr:    ":" + PORT,
	}

	// listen and serve
	fmt.Printf("serving on port :%v\n", PORT)
	log.Fatal(s.ListenAndServe())
}
