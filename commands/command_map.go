package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"pokedex/api"
)

func CommandMap(cfg *api.Config) error {
	poke_location, err := cfg.Get_pokemon_location(cfg.Next)
	if err != nil {
		fmt.Println("Couldn't load pokemon locations")
		return err
	}
    response := &api.LocationResponse{}
    json.Unmarshal(poke_location, response)


	cfg.Next = response.Next
	cfg.Prev = response.Previous

    fmt.Printf("\nLocations: \n")
	for _, location := range response.Results {
        fmt.Printf(" - %v\n", location.Name)
	}
    fmt.Println()

	return nil
}

func CommandMapB(cfg *api.Config) error {
	if cfg.Prev == nil {
		return errors.New("Can't go back! Your at the end of the list...")
    }

	poke_location, err := cfg.Get_pokemon_location(cfg.Prev)

    if err != nil {
        fmt.Println("Couldn't load pokemon locations")
        return err
    }

    response := &api.LocationResponse{}
    json.Unmarshal(poke_location, response)

	cfg.Next = response.Next
	cfg.Prev = response.Previous

    fmt.Printf("\nLocations: \n")
	for _, location := range response.Results {
        fmt.Printf(" - %v\n", location.Name)
	}
    fmt.Println()

	return nil
}
