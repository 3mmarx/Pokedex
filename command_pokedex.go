package main

import (
	"errors"
	"fmt"
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
