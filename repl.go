package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/api"
	"pokedex/commands"
	"strings"
)

func Start_repl(cfg *api.Config) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("welcome to pokedex! ;)")
	fmt.Println("\nUsage: ")
	commands.CommandHelp(cfg)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

        words := cleanInput(reader.Text())
		if len(words) == 0 {
            continue
        }

		command_name := words[0]
		command, exists := commands.GetCommands()[command_name]

        if exists {
            if len(words) == 2 {
                if command.Name == "explore" {
                    cfg.City = words[1]
                    err := command.Callback(cfg)
                    if err != nil {
                        fmt.Println(err)
                    }
                } else if command.Name == "catch" {
                    cfg.PokemonName = words[1]
                    err := command.Callback(cfg)
                    if err != nil {
                        fmt.Println(err)
                    }
                } else if command.Name == "inspect" {
                    cfg.PokemonName = words[1]
                    err := command.Callback(cfg)
                    if err != nil {
                        fmt.Println(err)
                    }
                }
            } else {
                err := command.Callback(cfg)
                if err != nil {
                    fmt.Println(err)
                }
            }
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
