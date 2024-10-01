package main

import (
	"os"

	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func commandExit(cfg *pokeapi.Config, args []string) error {
	os.Exit(0)
	return nil
}