package api

import (
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (cfg *Config) Get_pokemon_location(pageURL *string) ([]byte, error) {
	url := baseURL + "/location-area/?limit=20"
    if pageURL != nil {
        url = *pageURL
    }

    cache_url, found := cfg.Cache.Get(url)
    if !found {
        res, err := http.Get(url)
        if err != nil {
            return nil, err
        }
        defer res.Body.Close()

        body, err := io.ReadAll(res.Body)
        if err != nil {
            return nil, err
        }

        cfg.Cache.Add(url, body)
        return body, nil
    }
    return cache_url, nil
}
