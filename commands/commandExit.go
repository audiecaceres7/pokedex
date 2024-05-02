package commands

import (
	"os"
	"pokedex/api"
)

func CommandExit(cfg *api.Config) error {
    os.Exit(1)
    return nil
}
