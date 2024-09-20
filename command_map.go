package main

import (
	"fmt"
	"github.com/nukestarmaster/Pokedex.pokeapi"
)

func commandMap() error {
	fmt.Println("Map goes here")
	locations, err := ListLocationAreas(false)

}

func commandMapB() error {
	fmt.Println("Previous map goes here")
	return nil
}