package main

import (
	"fmt"

	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func commandMap(cfg *pokeapi.Config, args []string) error {
	loc, err := cfg.ListLocationAreas(false)
	if err != nil {
		return err
	}

	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapB(cfg *pokeapi.Config, args []string) error {
	loc, err := cfg.ListLocationAreas(true)
	if err != nil {
		return err
	}

	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}
	return nil
}