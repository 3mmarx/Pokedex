package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/3mmarx/pokedexcli/internal/pokeapi"
)

func catchPokemon(cfg *config, pokemon pokeapi.Pokemon) {
	// chance of caughting a strong pokemon is 20% or less and weak one is 90% or more.
	catchChance := 90 - (pokemon.BaseExperience*70)/608 // 608 is the max base experience for a pokemon
	randNum := rand.IntN(100)

	if randNum <= catchChance {
		fmt.Printf(Green+"%s was caught!\n"+Reset, pokemon.Name)
		cfg.caughtPokemons[pokemon.Name] = pokemon
	} else {
		fmt.Printf(Red+"%s escaped!\n"+Reset, pokemon.Name)
	}
}
