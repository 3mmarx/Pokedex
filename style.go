package main

import "github.com/3mmarx/pokedexcli/internal/pokeapi"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
var BrightRed = "\033[91m"
var BrightGreen = "\033[92m"
var BrightYellow = "\033[93m"
var BrightBlue = "\033[94m"
var BrightMagenta = "\033[95m"
var BrightCyan = "\033[96m"
var BrightWhite = "\033[97m"

func pokemonColor(pokemon pokeapi.Pokemon) string {
	total := 0

	for _, stat := range pokemon.Stats {
		total += stat.BaseStat
	}

	switch {
	case total >= 600:
		return BrightMagenta
	case total >= 500:
		return BrightYellow
	case total >= 400:
		return BrightCyan
	case total >= 300:
		return BrightGreen
	default:
		return Gray
	}

}
