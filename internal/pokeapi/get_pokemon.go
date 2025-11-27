package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name *string) (Pokemon, error) {
	url := baseURL + "/pokemon/1/"
	if name != nil {
		url = baseURL + "/pokemon/" + *name + "/"
	}

	val, ok := c.cache.Get(url)
	if ok {
		cachedData := Pokemon{}
		err := json.Unmarshal(val, &cachedData)
		if err != nil {
			return Pokemon{}, err
		}
		return cachedData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	locAreas := Pokemon{}
	err = json.Unmarshal(data, &locAreas)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	return locAreas, nil

}
