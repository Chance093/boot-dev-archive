package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Chance093/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(_ *state, _ command) error {
  const feedURL = "https://www.wagslane.dev/index.xml"
  ctx := context.Background()
  rssFeed, err := fetchFeed(ctx, feedURL)
  if err != nil {
    return err
  }

  fmt.Println(rssFeed)
  return nil
}

func handlerAddFeed(s *state, cmd command) error {
  if len(cmd.args) != 2 {
    return errors.New("provide name and url args in that order")
  }

  name := cmd.args[0]
  url := cmd.args[1]
  ctx := context.Background()

  user, err := s.db.GetUser(ctx, s.cfg.Current_user_name)
  if err != nil {
    return fmt.Errorf("could not get user: %v", err)
  }

  feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
    ID: uuid.New(),
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
    Name: name,
    Url: url,
    UserID: user.ID,
  })
  if err != nil {
    return fmt.Errorf("could not create feed: %v", err)
  }

  fmt.Printf("feed added: %v", feed)
  return nil
}

func handlerListFeeds(s *state, _ command) error {
  ctx := context.Background()
  feeds, err := s.db.GetAllFeeds(ctx)
  if err != nil {
    return fmt.Errorf("couldn't list feeds: %v", err)
  }

  if len(feeds) == 0 {
    fmt.Println("no feeds")
    return nil
  }

  for i, feed := range feeds {
    fmt.Printf("Feed %d:\n", i + 1)
    fmt.Printf("  * Name: %v\n", feed.Name)
    fmt.Printf("  * URL: %v\n", feed.Url)
    fmt.Printf("  * Username: %v\n", feed.Username)
    fmt.Println()
  }

  return nil
}
