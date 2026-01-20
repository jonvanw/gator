package main

import (
	"context"
	"fmt"

	"github.com/jonvanw/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	if len(cmd.args) != 1 {
		//return fmt.Errorf("fetch command requires exactly one argument, the URL.")
	} else { // TODO else condition should be moved outside of if/else block
		url = cmd.args[0]
	}

	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}