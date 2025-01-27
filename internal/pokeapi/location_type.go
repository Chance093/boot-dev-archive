package pokeapi

type ShallowLocationsResp struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
}

type ShallowLocationResp struct {
	Id                int                 `json:"id"`
	Name              string              `json:"name"`
	PokemonEncounters []PokemonInLocation `json:"pokemon_encounters"`
}

type PokemonInLocation struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}
