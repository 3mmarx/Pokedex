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

type PokemonLocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int           `json:"min_level"`
				MaxLevel        int           `json:"max_level"`
				ConditionValues []interface{} `json:"condition_values"`
				Chance          int           `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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
