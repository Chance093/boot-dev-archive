package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (r *RSSFeed) EscapeHTML() {
	r.Channel.Title = html.UnescapeString(r.Channel.Title)
	r.Channel.Description = html.UnescapeString(r.Channel.Description)
	for i := 0; i < len(r.Channel.Item); i++ {
		r.Channel.Item[i].Title = html.UnescapeString(r.Channel.Item[i].Title)
		r.Channel.Item[i].Description = html.UnescapeString(r.Channel.Item[i].Description)
	}
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
  client := &http.Client{
    Timeout: time.Second * 5,
  }

	r, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create http req: %v", err)
	}

	r.Header.Set("User-Agent", "gator")

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return nil, fmt.Errorf("failed to unmarshal xml: %v", err)
	}

  rssFeed.EscapeHTML()

	return &rssFeed, nil
}
