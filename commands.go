package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists available locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous available locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Takes an area and lists pokemon in that area",
			callback:    commandExplore,
		},
	}
}

func commandHelp(cfg *Config, args ...string) error {
	cmd := getCommands()
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range cmd {
		str := value.name + ": " + value.description
		fmt.Println(str)
	}
	fmt.Println("")
	return nil
}

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapF(cfg *Config, args ...string) error {
	res, err := cfg.Client.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *Config, args ...string) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	res, err := cfg.Client.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *Config, args ...string) error {
  if len(args) < 1 {
    return errors.New("Please give an area after the command")
  }

  area := args[0]

	pokemon, err := cfg.Client.GetPokemonList(area)
	if err != nil {
		return err
	}

	for _, pm := range pokemon {
		fmt.Println(pm)
	}

	return nil
}
