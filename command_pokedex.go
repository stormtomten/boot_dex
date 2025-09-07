package main

import (
	"fmt"
)

func commandPokedex(args []string, cfg *config) error {
	dex := cfg.Pokedex
	if len(dex) == 0 {
		return fmt.Errorf("No Pokemons cought!")
	}
	fmt.Println("Your Pokedex")
	for _, poke := range dex {
		fmt.Printf(" - %s\n", poke.Name)
	}
	return nil
}
