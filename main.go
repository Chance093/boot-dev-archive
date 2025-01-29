package main

import (
	"fmt"

	"github.com/Chance093/gator/internal/config"
)


func main() {
  cfg, err := config.Read()
  if err != nil {
    panic(err)
  }

  err = cfg.SetUser("chance")
  if err != nil {
    panic(err)
  }

  cfg, err = config.Read()
  if err != nil {
    panic(err)
  }

  fmt.Println(*cfg)
}
