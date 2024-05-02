package commands

import (
	"fmt"
	"pokedex/api"
)

func CommandPokedex(cfg *api.Config) error {
    if len(cfg.Pokedex) <= 0 {
        fmt.Println("Your pokedex is empty! Try catching some pokemon ;)")
    }
    
    fmt.Printf("\nYour pokedex ->\n")
	for key := range cfg.Pokedex {
        fmt.Printf(" - %v\n", key)
	}
    fmt.Println()
    return nil
}


