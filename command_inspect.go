package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No location area name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemons[pokemonName]

	if !ok {
		return errors.New("You haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	return nil
}
