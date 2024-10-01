package main

import (
	"fmt"

	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func commandInspect(cfg *pokeapi.Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon name found")
		return nil
	}

	pokemon, ok := cfg.Pokedex[args[0]]
	if !ok {
		fmt.Println(args[0], "not found in Pokedex. Please catch Pokemon before inspecting them.")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Base stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(" -", t.Type.Name)
	}

	return nil
}