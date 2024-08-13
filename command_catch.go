package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No location area name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(&pokemonName)

	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Trying to catch %s\nWith Base Experience: %v \n", pokemon.Name, pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("%s ran away!", pokemon.Name)
	}
	fmt.Printf("You caught %s!\n", pokemon.Name)
	cfg.caughtPokemons[pokemon.Name] = pokemon

	return nil
}
