package commands

import (
	"fmt"
	"pokedex/api"
	"strings"
)

func CommandPokedex(cfg *api.Config) error {
    if len(cfg.Pokedex) <= 0 {
        fmt.Println("Your pokedex is empty! Try catching some pokemon ;)")
    }

    maxLength := 0
    for key := range cfg.Pokedex {
        if len(key) > maxLength {
            maxLength = len(key)
        }
    }

    fmt.Println()
    fmt.Println("POKEDEX")
    fmt.Printf("╭────────────────%s────────╮\n", strings.Repeat("─", maxLength))
    for key := range cfg.Pokedex {
        fmt.Printf("- %v\n", key)
    }
    fmt.Printf("╰────────────────%s────────╯\n\n", strings.Repeat("─", maxLength))

    return nil
}


