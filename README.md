# Pokedex CLI

An interactive command-line Pokedex built in Go and powered by [PokéAPI](https://pokeapi.co/). Explore location areas, discover and catch Pokémon, and inspect the ones you've collected — all from a lightweight REPL.

API responses are cached in memory to keep repeated lookups fast and reduce unnecessary network requests.

## Features

- Interactive REPL with case-insensitive command parsing
- Browse paginated location areas
- Explore the Pokémon found in any location
- Catch Pokémon, with catch difficulty based on base experience
- Inspect caught Pokémon to view their stats and types
- List your collection with color-coded entries based on total stats
- Thread-safe, TTL-based in-memory cache for API responses

## Requirements

- [Go](https://go.dev/dl/) 1.27.3 or newer
- An internet connection to reach PokéAPI

## Installation

```bash
git clone https://github.com/3mmarx/Pokedex.git
cd Pokedex
go build -o pokedexcli
./pokedexcli
```

Or run it directly without building a binary:

```bash
go run .
```

## Usage

| Command | Description |
| --- | --- |
| `help` | Show the available commands |
| `map` | List the next 20 location areas |
| `mapb` | List the previous 20 location areas |
| `explore <location-area>` | List the Pokémon found in a location area |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | Show details for a Pokémon you've caught |
| `pokedex` | List every Pokémon you've caught |
| `exit` | Close the Pokedex |

### Example session

```text
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
- tentacool
- staryu
- magikarp
...

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  - hp: 20
  - attack: 10
  ...
Types:
  - water

Pokedex > pokedex
  - magikarp

Pokedex > exit
Closing the Pokedex... Goodbye!
```

## How It Works

- **REPL** — `repl.go` reads input from stdin, normalizes it to lowercase, and dispatches the command to its registered handler.
- **Commands** — each command lives in its own `command_*.go` file and is registered in `commands.go`.
- **API client** — `internal/pokeapi` wraps the PokéAPI endpoints for location areas and Pokémon, with a configurable HTTP timeout.
- **Cache** — `internal/pokecache` provides a mutex-protected in-memory cache. A background goroutine periodically removes expired entries.
- **Catching** — each Pokémon's `base_experience` determines its catch chance: stronger Pokémon are harder to catch, while weaker ones are easier.

## Project Structure

```text
.
├── main.go               # Entry point
├── repl.go               # REPL loop and input parsing
├── commands.go           # Command registry
├── command_*.go          # Individual command handlers
├── colors.go             # Terminal color constants
├── repl_test.go          # REPL input tests
└── internal/
    ├── pokeapi/          # PokéAPI client and response types
    └── pokecache/        # TTL-based in-memory cache
```

## Testing

```bash
go test ./...
```

## Acknowledgements

- Pokémon data provided by [PokéAPI](https://pokeapi.co/)
- Built as part of the [Boot.dev](https://www.boot.dev/) backend development path

Pokémon and Pokémon character names are trademarks of Nintendo, Game Freak, and Creatures Inc. This project is for educational purposes and is not affiliated with or endorsed by them.
