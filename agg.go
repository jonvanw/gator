package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jonvanw/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("fetch command requires exactly one argument, the fetch frequency.")
	} 

	fetchFrequency, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to parse fetch frequency: %v", err)
	}

	fmt.Printf("Starting fetch loop with frequency: %s\n", fetchFrequency)
	for {
		err := scrapeNextFeed(s, time.Now().Add(fetchFrequency))
		if err != nil {
			return fmt.Errorf("failed to scrape next feed: %v", err)
		}
		time.Sleep(fetchFrequency)
	}
}

func scrapeNextFeed(s *state, skipCuttoff time.Time) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background(), skipCuttoff)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("no pending feeds")
			return nil
		}
		return fmt.Errorf("failed to get next feed to fetch: %v", err)
	}

	fmt.Printf("Fetching feed: %s...\n", feed.Name)

	rss, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	fmt.Printf("Title: %v\n", rss.Channel.Title)
	for _, item := range rss.Channel.Item {
		fmt.Printf("   Item Title: %v\n", item.Title)
		fmt.Printf("   Item Link: %v\n", item.Link)
		fmt.Printf("   Item Description: %v\n", item.Description)
		fmt.Printf("   Item PubDate: %v\n", item.PubDate)
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %v", err)
	}

	fmt.Printf("Feed marked as fetched at time: %v\n", time.Now())

	return nil
}