package main

import (
	"time"

	"github.com/3mmarx/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		client:   pokeapi.NewClient(5*time.Minute, 5*time.Second),
		previous: nil,
		next:     nil,
	}

	startRepl(cfg)
}
