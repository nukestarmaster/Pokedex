package main

import (
	"fmt"
	"math/rand"

	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func commandCatch(cfg *pokeapi.Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon name found")
		return nil
	}

	pokemon, err := cfg.ReqPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Throwing a pokeball at", pokemon.Name, "...")
	catch_rate := 256 - pokemon.BaseExperience
	if rand.Intn(256) < catch_rate {
		fmt.Println(pokemon.Name, "was caught!")
		_, recorded := cfg.Pokedex[args[0]]
		if !recorded {
			cfg.Pokedex[args[0]] = pokemon
			fmt.Println(pokemon.Name, "was added to Pokedex!")
		}
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}
	return nil
}