package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		fmt.Println("feeds command does not require any arguments, arguments ignored.")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	fmt.Printf("Feeds:\n")
	for _, feed := range feeds {	
		fmt.Printf("  Name: %s\n", feed.Name)
		fmt.Printf("  URL: %s\n", feed.Url)
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("failed to get user info for feed: %v", err)
		}
		fmt.Printf("  User: %s\n", user.Name)	
	}
	return nil
}