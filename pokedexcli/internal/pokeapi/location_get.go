package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	if val, ok := c.cache.Get(url); ok {
		location := Location{}
		if err := json.Unmarshal(val, &location); err != nil {
			return Location{}, err
		}
		return location, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}
	location := Location{}
	if err = json.Unmarshal(data, &location); err != nil {
		return Location{}, err
	}
	c.cache.Add(url, data)
	return location, nil
}
