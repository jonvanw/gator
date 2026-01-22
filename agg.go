package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jonvanw/gator/internal/database"
	"github.com/jonvanw/gator/internal/rss"
	"github.com/lib/pq"
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
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			// Try alternative format
			publishedAt, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				publishedAt = time.Now()
			}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID: feed.ID,
		})
		if err != nil {
			// Check if error is a unique constraint violation on URL
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				// Unique constraint violation - ignore it
				continue
			}
			// Check error message as fallback
			if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate key") {
				// URL already exists - ignore it
				continue
			}
			return fmt.Errorf("failed to create post: %v", err)
		}
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %v", err)
	}

	fmt.Printf("Feed fetched completed at time: %v\n", time.Now())

	return nil
}