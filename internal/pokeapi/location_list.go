package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	if resp.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("response failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locAreas := LocationAreas{}
	err = json.Unmarshal(data, &locAreas)
	if err != nil {
		return LocationAreas{}, err
	}

	return locAreas, nil
}
