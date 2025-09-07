package main

import (
	"errors"
	"fmt"
)

func commandExplore(args []string, cfg *config) error {
	if len(args) != 1 {
		return errors.New("usage: explore <area_name>")
	}
	name := args[0]
	resp, err := cfg.Client.ExploreLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring: %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, e := range resp.Encounters {
		fmt.Printf(" - %s\n", e.Pokemon.Name)
	}
	return nil
}
