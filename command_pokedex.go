package main

import (
	"fmt"
)

func commandPokedex(conf *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, val := range conf.pokedex {
		fmt.Printf(" - %v\n", val.Name)
	}
	return nil
}
