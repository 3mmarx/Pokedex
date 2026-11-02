package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/3mmarx/pokedexcli/internal/pokecache"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	http.Client
	cache *pokecache.Cache
}

func NewClient(cacheInterval, timeout time.Duration) Client {
	cache := pokecache.NewCache(cacheInterval)
	return Client{
		cache:   cache,
		Timeout: timeout,
	}
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c Client) FetchLocationAreas(url *string) (LocationAreas, error) {
	defaultURL := BaseURL + "/location-area"
	targetURL := defaultURL
	if url != nil {
		targetURL = *url
	}

	cachedRes, ok := c.cache.Get(targetURL)
	if ok {
		var locationAreas LocationAreas
		if err := json.Unmarshal(cachedRes, &locationAreas); err != nil {
			return LocationAreas{}, err
		}

		return locationAreas, nil
	}

	res, err := c.Get(targetURL)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreas{}, fmt.Errorf("unexpected status: %s", res.Status)
	}

	var locationAreas LocationAreas

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	if err := json.Unmarshal(resBytes, &locationAreas); err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(targetURL, resBytes)

	return locationAreas, nil
}
