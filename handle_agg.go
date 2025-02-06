package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Chance093/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
  if len(cmd.args) != 1 {
    return errors.New("must provide time duration after command")
  }

  timeDuration := cmd.args[0]
  timeBetwixTicks, err := time.ParseDuration(timeDuration)
  if err != nil {
    return fmt.Errorf("Must use correct duration format: %v", err)
  }

  ticker := time.NewTicker(timeBetwixTicks)
  for ; ; <-ticker.C {
    if err := scrapeFeeds(s); err != nil {
      return err
    }
  }
}

func scrapeFeeds(s *state) error {
  ctx := context.Background()

  feed, err := s.db.GetNextFeedToFetch(ctx)
  if err != nil {
    fmt.Printf("could not get next feed to fetch: %v", err)
  }

  rssFeed, err := fetchFeed(ctx, feed.Url)
  if err != nil {
    return err
  }
  
  err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
    ID: feed.ID,
    LastFetchedAt: sql.NullTime{
      Time: time.Now().UTC(),
      Valid: true,
    },
  })

  printFeed(rssFeed)
  return nil
}

func printFeed(rssFeed *RSSFeed) {
  fmt.Printf("Feed %v fetched:\n", rssFeed.Channel.Title)
  for _, item:= range rssFeed.Channel.Item {
    fmt.Printf("  * %v\n", item.Title)
  }
  fmt.Println()
}
