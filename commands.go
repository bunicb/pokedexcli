package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

type config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type Area struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var commands map[string]cliCommand = make(map[string]cliCommand)

func init() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
		config:	  	 &config{},
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
		config:	  	 &config{},
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Display 20 locations areas",
		callback:    commandMap,
		config:	  	 &config{},
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Display previous 20 locations areas",
		callback:    commandMapb,
		config:	  	 &config{},
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

func commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	if commands["map"].config.Next != "" {
		url = commands["map"].config.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("failed to fetch data: status code %d", res.StatusCode)
	}
	var locations Area

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
	}

	commands["map"].config.Next = locations.Next
	commands["map"].config.Previous = locations.Previous

	for r := range locations.Results {
		fmt.Println(locations.Results[r].Name)
	}
	return nil
}

func commandMapb() error {
	if commands["map"].config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := commands["map"].config.Previous
	
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("failed to fetch data: status code %d", res.StatusCode)
	}
	var locations Area

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
	}

	commands["map"].config.Next = locations.Next
	commands["map"].config.Previous = locations.Previous

	for r := range locations.Results {
		fmt.Println(locations.Results[r].Name)
	}
	return nil
}