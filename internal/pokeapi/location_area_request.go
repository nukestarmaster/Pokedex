package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

func (c *Client) ListLocationAreas(back bool) (LocationAreaResp, error) {
	endpoint := "/location-area"
	
	fullURL := baseURL + endpoint
	var locations LocationAreaResp

	if !back && next_areas != "" {
		fullURL = next_areas
	}

	if back && previous_areas != "" {
		fullURL = previous_areas
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return locations, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locations, err
	}

	json.Unmarshal(data, &locations)
	next_areas = locations.Next
	previous_areas = locations.Previous
	return locations, nil
}