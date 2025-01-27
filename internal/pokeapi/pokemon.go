package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemonList(area string) ([]string, error) {
  res, err := c.getLocation(area)
  if err != nil {
    return nil, err
  }

  pokemon := []string{}
  for _, encounter := range res.PokemonEncounters {
    pokemon = append(pokemon, encounter.Pokemon.Name)
  }

  return pokemon, nil
}

func (c *Client) GetPokemon(name string) (*ShallowPokemonResp, error) {
	var pokemonResponse ShallowPokemonResp

  url := c.BaseUrl + "/pokemon/" + name

  // check for cached response
  val, exists := c.Cache.Get(url)
  if exists {
    if err := json.Unmarshal(val, &pokemonResponse); err != nil {
      return nil, errors.New("Error while decoding cached json")
    }

    return &pokemonResponse, nil
  }

  // request locations from api
	res, err := c.HTTPClient.Get(url)
	if err != nil {
    return nil, errors.New("Error while fetching pokemon")
	}
	defer res.Body.Close()

  if res.StatusCode == 404 {
		return nil, errors.New("Pokemon does not exist")
  }

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Not an ok status code")
	}

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, errors.New("Error while reading response body")
  }

	if err := json.Unmarshal(body, &pokemonResponse); err != nil {
		return nil, errors.New("Error while decoding json")
	}

  // store response in cache
  c.Cache.Add(url, body)

	return &pokemonResponse, nil
}
