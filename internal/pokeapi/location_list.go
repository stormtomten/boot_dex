package pokeapi

import (
	"encoding/json"
)

type ShallowLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(url string) (ShallowLocations, error) {
	var out ShallowLocations
	fullURL := url
	if fullURL == "" {
		fullURL = "location-area?offset=0&limit=20"
	}
	b, err := c.doGet(fullURL)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal(b, &out); err != nil {
		return out, err
	}

	return out, nil
}
