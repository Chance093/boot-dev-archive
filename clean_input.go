package main

import (
  "strings"
)

func cleanInput(input string) []string {
  lower := strings.ToLower(input)
  splitInput := strings.Fields(lower)
  return splitInput
}
