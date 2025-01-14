package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands(config *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(config) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists available locations",
			callback:    func() error { return commandMap(config, config.Next) },
		},
		"mapb": {
			name:        "map",
			description: "Lists previous available locations",
			callback:    func() error { return commandMap(config, config.Previous) },
		},
	}
}

func commandHelp(config *Config) error {
	cmd := getCommands(config)
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

func commandExit() error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *Config, url *string) error {
	res, err := getLocations(url)
	if err != nil {
		return err
	}

  config.Next = res.Next
  config.Previous = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}
