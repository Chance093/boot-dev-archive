package pokeapi

import (
	"net/http"
	"time"
)

type PokeClient struct {
	BaseUrl    string
	HTTPClient *http.Client
}

func NewClient(timeout time.Duration) *PokeClient {
	return &PokeClient{
		BaseUrl: "https://pokeapi.co/api/v2",
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}
