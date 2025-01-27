package pokeapi

import "errors"

type Pokedex struct {
  Pokedex map[string]ShallowPokemonResp
}

func NewPokedex() *Pokedex {
  return &Pokedex{
    Pokedex: map[string]ShallowPokemonResp{},
  }
}

func (p *Pokedex) AddPokemon(pm *ShallowPokemonResp) {
  _, exists := p.Pokedex[pm.Name]
  if exists {
    return 
  }

  p.Pokedex[pm.Name] = *pm
}

func (p *Pokedex) GetPokemonList() ([]string, error) {
  pokemonList := []string{}

  for key := range p.Pokedex {
    pokemonList = append(pokemonList, key)
  }

  if len(pokemonList) == 0 {
    return nil, errors.New("You have nothing in your pokedex")
  }

  return pokemonList, nil
}

func (p *Pokedex) GetPokemon(name string) (ShallowPokemonResp, error) {
  pokemon, exists := p.Pokedex[name]
  if !exists {
    return ShallowPokemonResp{}, errors.New("Pokemon does not exist in Pokedex")
  }

  return pokemon, nil
}
