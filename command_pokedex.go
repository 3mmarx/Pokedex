package main

import (
	"errors"
	"fmt"

	"github.com/3mmarx/pokedexcli/internal/pokeapi"
)

func CommandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemons) == 0 {
		return errors.New(Red + "Your pokedex is empty, try catching any pokemon." + Reset)
	}

	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf(pokemonColor(pokemon)+"  - %s\n"+Reset, pokemon.Name)
	}

	return nil
}

func pokemonColor(pokemon pokeapi.Pokemon) string {
	total := 0

	for _, stat := range pokemon.Stats {
		total += stat.BaseStat
	}

	switch {
	case total >= 600:
		return BrightMagenta
	case total >= 500:
		return BrightYellow
	case total >= 400:
		return BrightCyan
	case total >= 300:
		return BrightGreen
	default:
		return Gray
	}

}
