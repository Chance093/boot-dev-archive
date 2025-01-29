package main

import (
	"os"

	"github.com/Chance093/gator/internal/config"
)


func main() {
  const configFile = "/.gatorconfig.json"

  path, err := os.UserHomeDir()
  if err != nil {
    panic(err)
  }

  _, err = config.Read(path + configFile)
  if err != nil {
    panic(err)
  }
}

