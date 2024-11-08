package main

import (
  "strings"
)

func cleanInput(input string) []string {
  cleanedInput := strings.TrimSpace(input)
  splitInput := strings.Split(cleanedInput, " ")
  return splitInput
}
