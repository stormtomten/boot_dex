package main

type cliCommand struct {
	name        string
	description string
	callback    func([]string, *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Lists the location areas, 20 at a time, usage: map",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Steps back once in the 'map' command, usage: mapb",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a location area, usage: explore <area_name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon, usage: catch <name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects an owned Pokemon, usage: inspect <name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught Pokemon",
			callback:    commandPokedex,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
