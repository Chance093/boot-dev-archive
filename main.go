package main

import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
  Next *string
  Previous *string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
  config := Config{
    Next: nil,
    Previous: nil,
  }

	commands := getCommands(&config)

	fmt.Print("pokedex > ")

	for scanner.Scan() {
		txt := scanner.Text()
    cleanInput := cleanInput(txt)
    cmd, ok := commands[cleanInput[0]]
    if !ok {
      fmt.Println("Please enter a valid command (See help command).")
      fmt.Println("")
		  fmt.Print("pokedex > ")
      continue
    }

    err := cmd.callback()
    if err != nil {
      fmt.Println(err)
    }

		fmt.Print("pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
