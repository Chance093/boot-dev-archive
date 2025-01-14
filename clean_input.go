package main

import (
  "strings"
)

// Clean input and split into array of cmds
func cleanInput(input string) []string {
  lower := strings.ToLower(input)
  splitInput := strings.Fields(lower)
  return splitInput
}
