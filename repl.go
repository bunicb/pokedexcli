package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"

	"github.com/bunicb/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}


func startRepl(cfg *config) {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		input.Scan()

		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}
		
		userCommand := words[0]

		validCommand, ok := getCommands()[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := validCommand.callback(cfg)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}