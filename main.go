package main

import (
	"time"

	"github.com/stormtomten/boot_dex/internal/pokeapi"
)

func main() {
	cfg := &config{
		Client:  pokeapi.NewClient(5*time.Second, 5*time.Minute),
		Pokedex: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
