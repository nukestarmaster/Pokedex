package main

import (
	"fmt"

	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func commandHelp(cfg *pokeapi.Config, args []string) error {
	commands := getCommands()
	fmt.Println("Showing Help Menu")
	for _, c := range commands {
		fmt.Println(c.name, ": ", c.description)
	}
	return nil
}