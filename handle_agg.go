package main

import (
	"context"
	"fmt"
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
