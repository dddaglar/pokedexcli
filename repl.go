package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dddaglar/pokedexcli/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := &config{next: 0, previous: 0}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		command := text[0]
		com, ok := getCommand()[command]
		if ok {
			err := com.callback(conf)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	ntext := strings.ToLower(text)
	res := strings.Fields(ntext)
	return res
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          int
	previous      int
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the map of Pokemons",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of the map",
			callback:    commandMapb,
		},
	}
}
