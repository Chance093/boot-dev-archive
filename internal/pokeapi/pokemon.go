package pokeapi

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
