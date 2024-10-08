package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No location area name provided")
	}
	locationAreaName := args[0]

	fmt.Printf("Exploring %s...\n", locationAreaName)
	locationArea, err := cfg.pokeapiClient.GetLocationArea(&locationAreaName)

	if err != nil {
		return err
	}

	fmt.Printf("Found Pokemon: %s\n", locationArea.Name)
	for _, area := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", area.Pokemon.Name)
	}

	return nil
}
