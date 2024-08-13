package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
