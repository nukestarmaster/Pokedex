package pokeapi

import (
	"time"

	"github.com/nukestarmaster/Pokedex/internal/pokecache"
)



type Config struct {
	PokeapiClient   Client
	Cache 			pokecache.Cache
	next_areas		string
	previous_areas	string
	Pokedex			map[string]Pokemon
}

func Init(t int) Config {
	return Config{
		PokeapiClient: 	NewClient(time.Duration(t) * time.Second),
		Cache:			pokecache.NewCache(t),
		Pokedex:		make(map[string]Pokemon),
	}
}
