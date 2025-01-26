package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
	}
}

func commandHelp(cfg *Config) error {
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

func commandExit(cfg *Config) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapF(cfg *Config) error {
	res, err := cfg.PokeClient.GetLocations(cfg.Next)
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

func commandMapB(cfg *Config) error {
  if cfg.Previous == nil {
    return errors.New("you're on the first page")
  }

	res, err := cfg.PokeClient.GetLocations(cfg.Previous)
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
