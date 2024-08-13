package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name *string) (Pokemon, error) {
	endpoint := "/pokemon/" + *name
	fullUrl := baseUrl + endpoint

	// check cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit")
		locationAreasResp := Pokemon{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return Pokemon{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	locationAreasResp := Pokemon{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return Pokemon{}, err
	}

	// add to cache
	c.cache.Add(fullUrl, data)

	return locationAreasResp, nil
}
