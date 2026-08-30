package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonAPI := pokemonAPI{}
	if err = json.Unmarshal(data, &pokemonAPI); err != nil {
		return Pokemon{}, err
	}
	pokemonSpecies, err := c.GetSpeciesURL(pokemonAPI.SpeciesSh.URL)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{
		Pokemon: pokemonAPI,
		Species: pokemonSpecies,
	}
	data, err = json.Marshal(pokemon)
	if err == nil {
		c.cache.Add(url, data)
	}
	return pokemon, nil
}
