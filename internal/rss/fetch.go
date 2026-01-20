package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"net/http"
)

func FetchFeed(ctx context.Context, url string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator/0.1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch feed, HTTP status: %s", res.Status)
	}

	feed := &RSSFeed{}
	err = xml.NewDecoder(res.Body).Decode(feed)
	if err != nil {
		return nil, fmt.Errorf("failed to decode feed: %v", err)
	}

	// Unescape HTML entities in Title and Description fields
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return feed, nil
}