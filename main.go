package main

import (
  "time"
  "github.com/Chance093/pokedexcli/internal/pokeapi"
)

type Config struct {
  PokeClient *pokeapi.PokeClient
  Next *string
  Previous *string
}

func main() {
  pokeClient := pokeapi.NewClient(3 * time.Second)
  cfg := &Config{
    PokeClient: pokeClient,
  }

  startRepl(cfg)
}
