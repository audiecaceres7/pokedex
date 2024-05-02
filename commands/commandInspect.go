package commands

import (
	"fmt"
	"pokedex/api"
)

func CommandInspect(cfg *api.Config) error {        
    pokemon, found := cfg.Pokedex[cfg.PokemonName]
    if !found {
        fmt.Println("Pokemon not found, Try catching it first! ;)")
    } else {
        fmt.Println()
        fmt.Printf("name: %v\n", pokemon.Name)     
        fmt.Printf("Height: %v\n", pokemon.Height)     
        fmt.Printf("weight: %v\n", pokemon.Weight)     
        fmt.Printf("Stats: \n")
        for _, val := range pokemon.Stats {
            fmt.Printf("  - %v: %v\n", val.Stat.Name, val.BaseStat)
        }
        fmt.Println("Types: ")
        for _, val := range pokemon.Types {
            fmt.Printf("  - %v\n", val.Type.Name)
        }
        fmt.Println()
    }

    return nil
}
