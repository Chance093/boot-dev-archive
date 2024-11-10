package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
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
