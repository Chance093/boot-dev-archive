package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
  registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, fn func(*state, command) error) {
  c.registeredCommands[name] = fn
}

func (c *commands) run(s *state, cmd command) error {
  fn, exists := c.registeredCommands[cmd.name]
  if !exists {
    return errors.New("command not found")
  }

  return fn(s, cmd)
}


