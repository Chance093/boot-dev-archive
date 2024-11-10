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

func getLocations() (LocationResponse, error) {
	url := "https://pokeapi.co/api/v2/location"
	res, err := http.Get(url)

	if err != nil {
		return LocationResponse{}, errors.New("Error while fetching location")
	}
	defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    return LocationResponse{}, errors.New("Not an ok status code")
  }

	var locationResponse LocationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationResponse); err != nil {
		return LocationResponse{}, errors.New("Error while decoding json")
	}

	return locationResponse, nil
}
