package main

import (
	"context"
	"fmt"

	"github.com/jonvanw/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("unfollow command requires exactly one argument, the feed URL.")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user info for current user: %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %v", err)
	}

	fmt.Printf("Feed unfollowed successfully:\n")
	fmt.Printf("  Feed: %s\n", feed.Name)
	fmt.Printf("  User: %s\n", user.Name)
	return nil
}