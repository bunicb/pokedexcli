package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand = make(map[string]cliCommand)

func init() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

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