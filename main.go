package main

import (
	"time"

	"github.com/four88/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	caughtPokemons  map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:  pokeapi.NewClient(time.Hour),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
