package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jonvanw/gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2
	if len(cmd.args) > 1 {
		fmt.Println("browse command accepts at most one optional argument, the number of posts to browse.")
	} else if len(cmd.args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("failed to parse browselimit: %v", err)
		}
	}


	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name: s.cfg.CurrentUserName,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to get posts: %v", err)
	}

	fmt.Printf("Posts:\n")
	for _, post := range posts {
		fmt.Printf("Feed:  %s\n", post.FeedName)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL:   %s\n", post.Url)
		fmt.Printf("Date:  %s\n", post.PublishedAt)
		fmt.Printf("Description:   %s\n", post.Description)
		fmt.Printf("\n")
	}

	return nil
}