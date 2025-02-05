package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Chance093/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
  if len(cmd.args) != 1 {
    return errors.New("command takes 1 arg which is url")
  }

  url := cmd.args[0]
  ctx := context.Background()

  user, err := s.db.GetUser(ctx, s.cfg.Current_user_name)
  if err != nil {
    return fmt.Errorf("user not found: %v", err)
  }

  feed, err := s.db.GetFeedByUrl(ctx, url)
  if err != nil {
    return fmt.Errorf("feed not found: %v", err)
  }

  feed_follow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
    ID: uuid.New(),
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
    FeedID: feed.ID,
    UserID: user.ID,
  })
  if err != nil {
    return fmt.Errorf("couldn't create feed follow: %v", err)
  }

  fmt.Printf("Feed: %v\n", feed_follow.FeedName)
  fmt.Printf("Username: %v\n", feed_follow.UserName)
  return nil
}

func handlerFollowing(s *state, _ command) error {
  ctx := context.Background()

  user, err := s.db.GetUser(ctx, s.cfg.Current_user_name)
  if err != nil {
    return fmt.Errorf("failed to get user: %v", err)
  }

  feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
  if err != nil {
    return fmt.Errorf("failed to get feeds for user: %v", err)
  }

  if len(feeds) == 0 {
    fmt.Println("you aren't following any feeds")
    return nil
  }

  fmt.Println("Feeds you are following:")
  for _, feed := range feeds {
    fmt.Printf("  * %v\n", feed.Name)
  }
  
  return nil
}
