package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageUrl *string) (*ShallowLocationsResp, error) {
	var locationResponse ShallowLocationsResp

  url := c.BaseUrl + "/location-area"
  if pageUrl != nil {
    url = *pageUrl
  }

  // check for cached response
  val, exists := c.Cache.Get(url)
  if exists {
    if err := json.Unmarshal(val, &locationResponse); err != nil {
      return nil, errors.New("Error while decoding cached json")
    }

    return &locationResponse, nil
  }

  // request locations from api
	res, err := c.HTTPClient.Get(url)
	if err != nil {
    return nil, errors.New("Error while fetching location")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Not an ok status code")
	}

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, errors.New("Error while reading response body")
  }

	if err := json.Unmarshal(body, &locationResponse); err != nil {
		return nil, errors.New("Error while decoding json")
	}

  // store response in cache
  c.Cache.Add(url, body)

	return &locationResponse, nil
}

func (c *Client) getLocation(area string) (*ShallowLocationResp, error) {
	var locationResponse ShallowLocationResp

  url := c.BaseUrl + "/location-area/" + area

  // check for cached response
  val, exists := c.Cache.Get(url)
  if exists {
    if err := json.Unmarshal(val, &locationResponse); err != nil {
      return nil, errors.New("Error while decoding cached json")
    }

    return &locationResponse, nil
  }

  // request locations from api
	res, err := c.HTTPClient.Get(url)
	if err != nil {
    return nil, errors.New("Error while fetching location")
	}
	defer res.Body.Close()

  if res.StatusCode == 404 {
		return nil, errors.New("Area does not exist")
  }

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Not an ok status code")
	}

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, errors.New("Error while reading response body")
  }

	if err := json.Unmarshal(body, &locationResponse); err != nil {
		return nil, errors.New("Error while decoding json")
	}

  // store response in cache
  c.Cache.Add(url, body)

	return &locationResponse, nil
}
