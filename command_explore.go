package main

import (
	"fmt"

	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func commandExplore(cfg *pokeapi.Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No area name found")
		return nil
	}

	fmt.Printf("Exploring %s ...\n", args[0])

	pokemon, err := cfg.ListAreaPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, p := range pokemon {
		fmt.Printf(" - %s\n", p.Name)
	}

	return nil
}