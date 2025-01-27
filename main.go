package main

import (
	"time"

	"github.com/Chance093/pokedexcli/internal/pokeapi"
)

type Config struct {
	Client *pokeapi.Client
	Next       *string
	Previous   *string
}

func main() {
  const timeout = time.Second * 3
  const cacheInterval = time.Minute * 1
	c := pokeapi.NewClient(timeout, cacheInterval)

	cfg := &Config{
		Client: c,
	}

	startRepl(cfg)
}
