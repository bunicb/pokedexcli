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
		
		fmt.Printf("Your command was: %s\n", words[0])
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}