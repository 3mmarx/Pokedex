package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
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

	// chance of caughting a strong pokemon is 20% or less and weak one is 90% or more.
	catchChance := 90 - (pokemon.BaseExperience*70)/608 // 608 is the max base experience for a pokemon
	randNum := rand.IntN(100)

	if randNum <= catchChance {
		fmt.Printf(Green+"%s was caught!\n"+Reset, pokemon.Name)
		cfg.caughtPokemons[pokemon.Name] = pokemon
	} else {
		fmt.Printf(Red+"%s escaped!\n"+Reset, pokemon.Name)
	}

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
