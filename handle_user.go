package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Chance093/gator/internal/database"
	"github.com/google/uuid"
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

func handlerRegister(s *state, cmd command) error {
  if len(cmd.args) != 1 {
    return errors.New("No user provided as argument")
  }

  username := cmd.args[0]
  ctx := context.Background()
  id := uuid.New()

  _, err := s.db.GetUser(ctx, username)
  if err == nil {
    return fmt.Errorf("user %v already exists", username)
  }

  _, err = s.db.CreateUser(ctx, database.CreateUserParams{
    ID: id,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
    Name: username,
  })
  if err != nil {
    return err
  }

  return nil
}
