package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/four88/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) ListLocationArea(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	// check cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// add to cache
	c.cache.Add(fullUrl, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(name *string) (LocationArea, error) {
	endpoint := "/location-area/" + *name
	fullUrl := baseUrl + endpoint

	// check cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit")
		locationAreasResp := LocationArea{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreasResp := LocationArea{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationArea{}, err
	}

	// add to cache
	c.cache.Add(fullUrl, data)

	return locationAreasResp, nil
}
