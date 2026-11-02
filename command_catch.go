package main

import (
	"errors"
	"fmt"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New(Red + "You must provide one pokemon." + Reset)
	}

	pokemon, err := cfg.client.FetchPokemonByName(args[0])
	if err != nil {
		return err
	}

	_, ok := cfg.caughtPokemons[pokemon.Name]
	if ok {
		fmt.Printf(Blue+"%s already caught!\n"+Reset, pokemon.Name)
		return nil
	}

	throwingPokeball(args[0])
	catchPokemon(cfg, pokemon)

	return nil
}

// adding some animation when throwing the pokeball
func throwingPokeball(name string) {
	fmt.Printf(Yellow+"Throwing a Pokeball at %s", name)

	for range 3 {
		for range 3 {
			time.Sleep(300 * time.Millisecond)
			fmt.Print(".")
		}

		time.Sleep(400 * time.Millisecond)
		fmt.Print("\b\b\b   \b\b\b")
	}

	fmt.Print(Reset + "\n")
}
