package main

import (
	"pokedex/api"
	"pokedex/pokecache"
	"time"
)

func main() {
	c := pokecache.New_cache(5 * time.Minute)
    cfg := &api.Config{
		Cache: c,
        Pokedex: make(map[string]api.PokemonResponse),
	}
	Start_repl(cfg)
}
