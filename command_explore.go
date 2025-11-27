package main

import (
	"fmt"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must provide a location name")
	}
	name := args[0]
	pokeListResp, err := conf.pokeapiClient.PokemonsInLocation(&name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v \n", pokeListResp.Name)
	fmt.Println("Found Pokemon:")
	for _, item := range pokeListResp.PokemonEncounters {
		fmt.Printf(" - %s\n", item.Pokemon.Name)
	}
	return nil
}
