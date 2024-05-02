package commands

import (
	"encoding/json"
	"fmt"
	"pokedex/api"
)

const (
    baseURL = "https://pokeapi.co/api/v2/";
)

func CommandExplore(cfg *api.Config) error {
    url := baseURL + "location-area/" + cfg.City
    fmt.Printf("\nExploring %v ...\n", cfg.City)
    res, err := cfg.Get_pokemon_location(&url)

    cfg.Cache.Add(url, res)
    response := &api.CityResponse{}
    json.Unmarshal(res, response)

    if err != nil {
        fmt.Print("Couldn't load response\n")
        return err
    }

    fmt.Println("\nFound Pokemon :")
    for _, val := range response.PokemonEncounters {
        fmt.Printf("  - %v\n", val.Pokemon.Name)
    }
    fmt.Println()

    return nil
}
