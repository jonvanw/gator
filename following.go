package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		fmt.Println("following command does not require any arguments, arguments ignored.")
	}

	userName := s.cfg.CurrentUserName

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("failed to get feed follows for user: %v", err)
	}

	fmt.Printf("Feed follows for current user (%s):\n", userName)
	for _, feedFollow := range feedFollows {
		fmt.Printf(" - %s\n", feedFollow.FeedName)
	}
	return nil
}