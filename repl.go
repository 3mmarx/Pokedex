package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/3mmarx/pokedexcli/internal/pokeapi"
)

type config struct {
	commands       map[string]cliCommand
	caughtPokemons map[string]pokeapi.Pokemon
	client         pokeapi.Client
	next           *string
	previous       *string
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error:", err)
			break
		}

		parts := cleanInput(scanner.Text())

		if len(parts) == 0 {
			continue
		}

		command, ok := cfg.commands[parts[0]]
		if !ok {
			fmt.Println(Red + "Unknown command" + Reset)
			continue
		}

		args := parts[1:]
		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(err)
		}
	}
}
