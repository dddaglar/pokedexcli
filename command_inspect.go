package main

import (
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	name := args[0]
	val, ok := conf.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", val.Name)
	fmt.Printf("Height: %v\n", val.Height)
	fmt.Printf("Weight: %v\n", val.Weight)

	fmt.Println("Stats:")
	for _, stat := range val.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, item := range val.Types {
		fmt.Printf(" -%v\n", item.Type.Name)
	}
	return nil
}
