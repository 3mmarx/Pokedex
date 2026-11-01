package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/3mmarx/pokedexcli/internal/pokeapi"
)

type config struct {
	commands map[string]cliCommand
	client   pokeapi.Client
	next     *string
	previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20.",
		callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 location areas.",
		callback:    commandMapb,
	}

	return commands
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

		input := scanner.Text()
		if len(input) == 0 {
			continue
		}

		command, ok := cfg.commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}
