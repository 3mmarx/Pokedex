package main

import "github.com/3mmarx/pokedexcli/internal/pokeapi"

func main() {
	cfg := &config{
		commands: getCommands(),
		client:   pokeapi.NewClient(),
		previous: nil,
		next:     nil,
	}

	startRepl(cfg)
}
