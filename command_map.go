package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationUrl = resp.Next
	cfg.prevLocationUrl = resp.Previous
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("Your are at the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationUrl = resp.Next
	cfg.prevLocationUrl = resp.Previous
	return nil
}
