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

func (c Client) FetchPokemonsByLocation(locationName string) (PokemonLocationArea, error) {
	url := BaseURL + "/location-area/" + locationName

	cachedRes, ok := c.cache.Get(url)
	if ok {
		var pokemons PokemonLocationArea
		if err := json.Unmarshal(cachedRes, &pokemons); err != nil {
			return PokemonLocationArea{}, err
		}

		return pokemons, nil
	}

	res, err := c.Get(url)
	if err != nil {
		return PokemonLocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PokemonLocationArea{}, fmt.Errorf("unexpected status: %s", res.Status)
	}

	var pokemons PokemonLocationArea

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonLocationArea{}, err
	}

	if err := json.Unmarshal(resBytes, &pokemons); err != nil {
		return PokemonLocationArea{}, err
	}

	c.cache.Add(url, resBytes)

	return pokemons, nil
}

func (c Client) FetchPokemonByName(pokemonName string) (Pokemon, error) {
	url := BaseURL + "/pokemon/" + pokemonName

	cachedRes, ok := c.cache.Get(url)
	if ok {
		var pokemon Pokemon
		if err := json.Unmarshal(cachedRes, &pokemon); err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	res, err := c.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status: %s", res.Status)
	}

	var pokemon Pokemon

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(resBytes, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, resBytes)

	return pokemon, nil
}
