package main

import (
	"fmt"

	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func commandPokedex(cfg *pokeapi.Config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.Pokedex {
		fmt.Println(" -", p.Name)
	}
	return nil
}