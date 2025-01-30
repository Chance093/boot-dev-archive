package main

import (
	"log"
	"os"

	"github.com/Chance093/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
    log.Fatal(err)
	}

  cmds := commands{
    registeredCommands: make(map[string]func(*state, command) error),
  }
  cmds.register("login", handlerLogin)

  s := &state{
    cfg,
  }

  rawArgs := os.Args[1:]
  if len(rawArgs) < 1 {
    log.Fatal(err)
  }

  cmdName, cmdArgs := rawArgs[0], rawArgs[1:]

  if err := cmds.run(s, command{cmdName, cmdArgs}); err != nil {
    log.Fatal(err)
  }
}
