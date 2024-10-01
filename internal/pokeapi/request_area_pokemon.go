package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Config) ListAreaPokemon(name string) ([]PokemonRef, error) {
	endpoint := "location-area/"

	fullURL := baseURL + endpoint + name

	var pokemon []PokemonRef

	data, ok := c.Cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return pokemon, err
		}

		resp, err := c.PokeapiClient.httpClient.Do(req)
		if err != nil {
			return pokemon, err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return pokemon, fmt.Errorf(resp.Status)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return pokemon, err
		}

		c.Cache.Add(fullURL, data)
	}

	var areaResp AreaPokemonResp
	json.Unmarshal(data, &areaResp)

	for _, p := range areaResp.PokemonEncounters {
		pokemon = append(pokemon, p.Pokemon)
	}

	return pokemon, nil
}
