package pokeapi

import "net/http"

func (c *Client) ListLocationAreas() (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		
	}
}