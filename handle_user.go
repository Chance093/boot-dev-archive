package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
  if len(cmd.args) != 1 {
    return errors.New("No user provided as argument")
  }
  
  username := cmd.args[0]

  if err := s.cfg.SetUser(username); err != nil {
    return err
  }

  fmt.Printf("User %v has been set\n", username)
  return nil
}
