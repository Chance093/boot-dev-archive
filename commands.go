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
	}
}

func commandHelp() error {
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

func commandExit() error {
	os.Exit(0)
	return nil
}

