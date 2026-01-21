package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jonvanw/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("follow command requires exactly one argument, the feed URL.")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user info for current user: %v", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	fmt.Printf("Feed follow created successfully:\n")
	fmt.Printf("  Follow ID: %s\n", feedFollow.ID)
	fmt.Printf("  Feed ID: %s\n", feedFollow.FeedID)
	fmt.Printf("  User ID: %s\n", feedFollow.UserID)
	return nil
}