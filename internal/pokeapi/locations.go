package pokeapi

import (
  "net/http"
  "errors"
  "encoding/json"
)

func (pc *PokeClient) GetLocations(pageUrl *string) (*ShallowLocationResp, error) {
  url := pc.BaseUrl + "/location-area"
  if pageUrl != nil {
    url = *pageUrl
  }

	res, err := pc.HTTPClient.Get(url)
	if err != nil {
    return nil, errors.New("Error while fetching location")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Not an ok status code")
	}

	var locationResponse ShallowLocationResp
	if err := json.NewDecoder(res.Body).Decode(&locationResponse); err != nil {
		return nil, errors.New("Error while decoding json")
	}

	return &locationResponse, nil
}
