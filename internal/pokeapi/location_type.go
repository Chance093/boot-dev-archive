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
Id                int `json:"id"`
Name              string `json:"name"`
PokemonEncounters []Pokemon `json:"pokemon_encounters"`
}

type Pokemon struct {
	Pokemon struct {
Name string `json:"name"`
} `json:"pokemon"`
}
