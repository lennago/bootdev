package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetSpecies(speciesName string) (PokemonSpecies, error) {
	url := baseURL + "/pokemon-species/" + speciesName
	return c.GetSpeciesURL(url)
}

func (c *Client) GetSpeciesURL(url string) (PokemonSpecies, error) {
	if val, ok := c.cache.Get(url); ok {
		species := PokemonSpecies{}
		if err := json.Unmarshal(val, &species); err != nil {
			return PokemonSpecies{}, err
		}
		return species, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonSpecies{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonSpecies{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonSpecies{}, err
	}
	species := PokemonSpecies{}
	if err = json.Unmarshal(data, &species); err != nil {
		return PokemonSpecies{}, err
	}
	c.cache.Add(url, data)
	return species, nil
}
