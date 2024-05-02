package commands

import "pokedex/api"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*api.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 locations",
			Callback:    CommandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Explores pokemon area by name",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon by name",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Insept pokemon from Pokedex",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Check your collection of Pokemon",
            Callback:    CommandPokedex,
		},
	}
}
