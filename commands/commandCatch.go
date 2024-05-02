package commands

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"pokedex/api"
	"time"
)

const (
    POKEMON_MAX_BASE_XP = 635
)

func CommandCatch(cfg *api.Config) error {
    url := baseURL + "pokemon/" + cfg.PokemonName
    res, err := cfg.Get_pokemon_location(&url)
    if err != nil {
        fmt.Println("Pokemon not Found...")
        return err
    }

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    count := 0
    fmt.Printf("Throwing a Pokeball at %v ", cfg.PokemonName)

    for {
        select {
        case <-ticker.C:
            fmt.Print(".")
            count++
        }

        if count == 3 {
            fmt.Print("\n")
            break
        }
    }

    pokemon := &api.PokemonResponse{}
    json.Unmarshal(res, pokemon)
   
    random_num, _ := rand.Int(rand.Reader, big.NewInt(POKEMON_MAX_BASE_XP))
    base_xp := big.NewInt(int64(pokemon.BaseExperience))

    if random_num.Cmp(base_xp) == +1 {
        cfg.Pokedex[pokemon.Name] = *pokemon;
        fmt.Printf("\n%v was caught!\n", pokemon.Name) 
        fmt.Println(`
                @@@@@@@&&&&&&@@@,
            @@@&(((((((((((((((((&@@@
        @@&       ((((((((((((((((((&@@    
       @@&         ((((((((((((((((((((&@   
      @&(         (((((((((((((((((((((((&@ 
    @@&((((   ,((((((((((########&((((((((&@
    @&(((((((((((((((((##&       ##((((((((&
    @(((((&#############&         #####&((((
    &%%##################&       ###########
    @%%%%&**            &########%     &####
     &*******                              &
     \&*******                            &
      *&********                         & 
       **&*********                   *&    
        ***&**************     ******&      
              *&*****************&*        
        `)
        fmt.Printf("\nNow you can Inspect your pokemon by using the 'inspect <pokemon-name>' command\n") 
        fmt.Println()
    } else {
        fmt.Printf("\n%v escaped!\n", pokemon.Name) 
        fmt.Println()
    }
    return nil
}
