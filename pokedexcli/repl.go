package main

import (
	"bufio"
	"fmt"
	"os"
)


func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	fmt.Print("pokedex > ")

	for scanner.Scan() {
		txt := scanner.Text()
    cleanInput := cleanInput(txt)
    command := cleanInput[0]
    args := cleanInput[1:]

    cmd, exists := commands[command]
    if !exists {
      fmt.Println("Please enter a valid command (See help command).")
      fmt.Println("")
		  fmt.Print("pokedex > ")
      continue
    }

    err := cmd.callback(cfg, args...)
    if err != nil {
      fmt.Println(err)
    }

		fmt.Print("pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
