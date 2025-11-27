package main

import (
	"fmt"
	"math/rand"
)

func isCatched(baseExp int) bool {
	if baseExp > 400 {
		baseExp = 400
	}
	if baseExp < 20 {
		baseExp = 20
	}
	random := rand.Intn(400)
	if random > baseExp {
		return true
	} else {
		return false
	}

}

func commandCatch(conf *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	name := args[0]
	val, ok := conf.pokedex[name]
	if ok {
		fmt.Printf("%v is already in your Pokedex \n", val.Name)
		return nil
	}
	pokeDetailsResp, err := conf.pokeapiClient.GetPokemon(&name)
	if err != nil {
		return err
	}
	pName := pokeDetailsResp.Name
	fmt.Printf("Throwing a Pokeball at %v... \n", pName)
	if isCatched(pokeDetailsResp.BaseExperience) {
		fmt.Printf("%v was caught!\n", pName)
		conf.pokedex[pName] = pokeDetailsResp
	} else {
		fmt.Printf("%v escaped!\n", pName)
	}
	return nil
}
