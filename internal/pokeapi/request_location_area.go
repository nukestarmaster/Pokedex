package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Config) ListLocationAreas(back bool) (LocationAreaResp, error) {
	endpoint := "location-area?offset=0&limit=20"
	
	fullURL := baseURL + endpoint
	var locations LocationAreaResp

	if !back && c.next_areas != "" {
		fullURL = c.next_areas
	}

	if back && c.previous_areas != "" {
		fullURL = c.previous_areas
	}

	data, ok := c.Cache.Get(fullURL)
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return locations, err
		}

		resp, err := c.PokeapiClient.httpClient.Do(req)
		if err != nil {
			return locations, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return locations, err
		}

		c.Cache.Add(fullURL, data)
	}

	json.Unmarshal(data, &locations)
	c.next_areas = locations.Next
	c.previous_areas = locations.Previous
	return locations, nil
}