package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(args []string, cfg *config) error {
	if len(args) != 1 {
		return errors.New("usage: catch <name>")
	}
	name := args[0]
	pokemon, err := cfg.Client.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	res := rand.Intn(pokemon.Base_experience)
	if res > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}
	fmt.Printf("%s was caught!\n", name)
	cfg.Pokedex[pokemon.Name] = pokemon
	fmt.Println("You may now inpect it with the inspect command.")

	return nil
}
