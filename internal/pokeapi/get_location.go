package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonsInLocation(locArea *string) (PokemonList, error) {
	url := baseURL + "/location-area/1/"
	if locArea != nil {
		url = baseURL + "/location-area/" + *locArea + "/"
	}
	//put cache somewhere here
	val, ok := c.cache.Get(url)
	if ok {
		cachedData := PokemonList{}
		err := json.Unmarshal(val, &cachedData)
		if err != nil {
			return PokemonList{}, err
		}
		return cachedData, nil
	}

	//hit HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonList{}, err
	}
	if resp.StatusCode > 299 {
		return PokemonList{}, fmt.Errorf("response failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonList{}, err
	}

	Pokemons := PokemonList{}
	err = json.Unmarshal(data, &Pokemons)
	if err != nil {
		return PokemonList{}, err
	}
	c.cache.Add(url, data)
	return Pokemons, nil

}
