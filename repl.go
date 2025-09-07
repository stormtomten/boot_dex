package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/stormtomten/boot_dex/internal/pokeapi"
)

type config struct {
	Client   *pokeapi.Client
	Next     string
	Previous string
	Pokedex  map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		words := cleanInput(scan.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(words[1:], cfg)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
