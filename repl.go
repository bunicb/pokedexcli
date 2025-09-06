package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		input.Scan()

		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}
		
		userCommand := words[0]

		validCommand, ok := commands[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := validCommand.callback()
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