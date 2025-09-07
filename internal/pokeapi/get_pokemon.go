package pokeapi

import (
	"encoding/json"
	"errors"
	"strings"
)

type Pokemon struct {
	Name            string `json:"name"`
	Base_experience int    `json:"base_experience"`
	Height          int    `json:"height"`
	Weight          int    `json:"weight"`
	Stats           []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var out Pokemon
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return out, errors.New("invalid name")
	}
	path := "pokemon/" + name

	b, err := c.doGet(path)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal(b, &out); err != nil {
		return out, err
	}

	return out, nil
}
