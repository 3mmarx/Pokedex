package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	http.Client
}

func NewClient() Client {
	return Client{}
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

	res, err := c.Get(targetURL)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	var locationAreas LocationAreas
	if err := json.NewDecoder(res.Body).Decode(&locationAreas); err != nil {
		return LocationAreas{}, err
	}

	fmt.Println(locationAreas)

	return locationAreas, nil
}
