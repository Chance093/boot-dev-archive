package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type LocationResponse struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
}

func getLocations(url *string) (*LocationResponse, error) {
  if url == nil {
    defaultUrl:= "https://pokeapi.co/api/v2/location"
    url = &defaultUrl
  }

	res, err := http.Get(*url)
	if err != nil {
		return nil, errors.New("Error while fetching location")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Not an ok status code")
	}

	var locationResponse LocationResponse
	if err := json.NewDecoder(res.Body).Decode(&locationResponse); err != nil {
		return nil, errors.New("Error while decoding json")
	}

	return &locationResponse, nil
}
