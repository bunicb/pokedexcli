package main

import (
	"fmt"
	"error"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	location, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = location.Next
	cfg.prevLocationsURL = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}