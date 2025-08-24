package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		if !input.Scan() {
			break
		}
		line := input.Text()
		words := cleanInput(line)
		fmt.Println("Your command was:", words[0])
	}
}