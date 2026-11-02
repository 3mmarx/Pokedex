package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New(Red + "You must provide one location area." + Reset)
	}

	pokemons, err := cfg.client.FetchPokemonsByLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf(Yellow+"Exploring %s...\n"+Reset, args[0])
	for _, p := range pokemons.PokemonEncounters {
		fmt.Printf(Gray+"- %s\n"+Reset, p.Pokemon.Name)
	}

	return nil
}
