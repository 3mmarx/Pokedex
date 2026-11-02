package main

import (
	"time"

	"github.com/3mmarx/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:       getCommands(),
		caughtPokemons: map[string]pokeapi.Pokemon{},
		client:         pokeapi.NewClient(5*time.Minute, 5*time.Second),
		previous:       nil,
		next:           nil,
	}

	startRepl(cfg)
}
