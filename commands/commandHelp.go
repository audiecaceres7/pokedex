package commands

import (
	"fmt"
	"pokedex/api"
)

func CommandHelp(cfg *api.Config) error {
    fmt.Println()
    for key, value := range GetCommands() {
        fmt.Printf("  %v: %v\n", key, value.Description)
    }
    fmt.Println()
    return nil
}
