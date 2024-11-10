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
    "map": {
      name: "map",
      description: "Lists available locations",
      callback: commandMap,
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

func commandMap() error {
  res, err := getLocations()
  if err != nil {
    return err
  }

  for _, location := range res.Results {
    fmt.Println(location.Name)
  }

  return nil
}

