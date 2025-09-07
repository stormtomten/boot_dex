package pokeapi

import (
	"encoding/json"
	"errors"
	"strings"
)

type LocationInfo struct {
	/*
		Location  []struct {
			Name: string	`json:"name"`
			Url: string 	`json:"url"`
		} `json:"location"`
	*/
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ExploreLocation(name string) (LocationInfo, error) {
	var out LocationInfo
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return out, errors.New("invalid location")
	}
	path := "location-area/" + name
	b, err := c.doGet(path)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal(b, &out); err != nil {
		return out, err
	}

	return out, nil
}
