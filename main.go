package main

import (
	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

const t = 300

func main() {
	cfg := pokeapi.Init(t)
	Repl(&cfg)
}

