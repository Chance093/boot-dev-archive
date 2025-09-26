package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Chance093/gator/internal/database"
	"github.com/google/uuid"
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
		return fmt.Errorf("could not get next feed to fetch: %v", err)
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})

  if err := addPostsToDB(s, ctx, rssFeed, feed); err != nil {
    return fmt.Errorf("error while adding posts to db: %v", err)
  }
	return nil
}

func addPostsToDB(s *state, ctx context.Context, rssFeed *RSSFeed, feed database.Feed) error {
	for _, item := range rssFeed.Channel.Item {
		pubTime, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", item.PubDate)
		if err != nil {
      pubTimeAgain, err := time.Parse("Mon, 02 Jan 2006 15:04:05 Z0700", item.PubDate)
      if err != nil {
        return fmt.Errorf("could not parse pubDate %v: %v", item.PubDate, err)
      }
      pubTime = pubTimeAgain
		}

    post, err := s.db.CreatePost(ctx, database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  pubTime,
				Valid: true,
			},
			FeedID: feed.ID,
		})
    if err != nil {
      if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
        fmt.Printf("post '%v' already exists\n", item.Title)
        continue
      }

			return fmt.Errorf("could not save post to db: %v\n", err)
    }

    fmt.Printf("post saved: %v\n", post.Title)
    fmt.Println()
	}
  return nil
}
