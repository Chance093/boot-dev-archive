package main

import (
	"context"
	"fmt"

	"github.com/Chance093/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
  return func(s *state, cmd command) error {
    ctx := context.Background()

    user, err := s.db.GetUser(ctx, s.cfg.Current_user_name)
    if err != nil {
      return fmt.Errorf("user not found: %v", err)
    }

    return handler(s, cmd, user)
  }
}
