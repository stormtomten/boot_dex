package main

import (
	"errors"
	"fmt"
)

func commandMapf(args []string, cfg *config) error {
	if len(args) != 0 {
		return errors.New("Unexpected arguments")
	}
	url := cfg.Next
	resp, err := cfg.Client.ListLocationAreas(url)
	if err != nil {
		return err
	}
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil
}

func commandMapb(args []string, cfg *config) error {
	if len(args) != 0 {
		return errors.New("Unexpected arguments")
	}
	url := cfg.Previous
	if url == "" {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.Client.ListLocationAreas(url)
	if err != nil {
		return err
	}
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	return nil
}
