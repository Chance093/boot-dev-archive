package main

import (
	"time"

	"github.com/Chance093/pokedexcli/internal/pokeapi"
)

type Config struct {
	Client   *pokeapi.Client
	Next     *string
	Previous *string
  Pokedex  *pokeapi.Pokedex
}

func main() {
	const timeout = time.Second * 3
	const cacheInterval = time.Minute * 1
	c := pokeapi.NewClient(timeout, cacheInterval)
  p := pokeapi.NewPokedex()

	cfg := &Config{
		Client: c,
    Pokedex: p,
	}

	startRepl(cfg)
}
