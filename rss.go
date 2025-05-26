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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create new GET req with NewRequestWithContext
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	// Set Header to gator
	req.Header.Set("User-Agent", "gator")

	// Perf HTTP req using http.CLient for more control (set timeout etc.title)
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing get request: %w", err)
	}
	// close response body to avoid resource leak
	defer res.Body.Close()

	// check if response status code is 200 OK
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status response: %s", res.Status)
	}

	// io.ReadAll(res.Body) // Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %w", err)
	}

	decodeFeed(&feed)

	return &feed, nil
}

func decodeFeed(feed *RSSFeed) {
	// Decode HTML entititins for RSS feed
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	// Decode each RSSitem in the feed, using i, _ instead of _,item because the latter makes a copy of item, not editing the &rssFeed
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
}
