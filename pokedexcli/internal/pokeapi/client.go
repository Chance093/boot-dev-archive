package pokeapi

import (
	"net/http"
	"time"

	"github.com/Chance093/pokedexcli/internal/pokecache"
)

type Client struct {
	BaseUrl    string
	HTTPClient *http.Client
	Cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		BaseUrl: "https://pokeapi.co/api/v2",
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(cacheInterval),
	}
}
