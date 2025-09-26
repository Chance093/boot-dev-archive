package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
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
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemon in your pokedex",
			callback:    commandPokedex,
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

func commandCatch(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Please give a pokemon you wish to catch")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	pokemon, err := cfg.Client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	isCaught := isCatch(pokemon.BaseExperience)
	if isCaught {
		cfg.Pokedex.AddPokemon(pokemon)
		fmt.Printf("%v was caught!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%v escaped!\n", pokemon.Name)
	return nil
}

func isCatch(base_xp int) bool {
	flooredInt := int(math.Floor(float64(base_xp / 20)))
	random := rand.Intn(flooredInt)
	guess := rand.Intn(flooredInt)

	if guess == random {
		return true
	}

	return false
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Please give a pokemon you wish to inspect")
	}

	pokemonName := args[0]

	stats, err := cfg.Pokedex.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %v\n", stats.Name)
	fmt.Printf("Height: %v\n", stats.Height)
	fmt.Printf("Weight: %v\n", stats.Weight)
	fmt.Println("Stats:")
	for _, stat := range stats.Stats {
		fmt.Printf("  - %v: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ptype := range stats.Types {
		fmt.Printf("  - %v\n", ptype.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *Config, _ ...string) error {
	pokemon, err := cfg.Pokedex.GetPokemonList()
	if err != nil {
		return err
	}

	fmt.Println("Your Pokedex:")
	for _, pm := range pokemon {
		fmt.Printf("  - %v\n", pm)
	}

	return nil
}
