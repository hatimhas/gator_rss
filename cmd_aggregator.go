package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("error: not enough arguments provided for aggregator command")
	}
	timeArgs := cmd.arguments[0]
	timeBetweenReq, err := time.ParseDuration(timeArgs)
	if err != nil {
		return fmt.Errorf("error parsing time duration: %w", err)
	}
	log.Printf("Collecting feeds every %s...", timeBetweenReq)
	ticker := time.NewTicker(timeBetweenReq)
	for ; ; <-ticker.C {
		scrapeFeed(s)
	}
}

// Getting the next feed: You need to query your database to find the feed that's most overdue for fetching.
// Marking it fetched: Once you've selected a feed, you need to update its record in the database to show that it's just been fetched, so you don't immediately pick it again.
// Fetching the feed: You'll use the fetchFeed function, just like you did before, but this time you'll use the URL retrieved from the database.
// Processing the posts: You'll loop through the items (posts) from the fetched feed and, for now, just print their titles.
func scrapeFeed(s *state) error {
	nextFeedURL, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	nextFeedData, err := s.db.GetFeedByURL(context.Background(), nextFeedURL)
	if err != nil {
		return fmt.Errorf("error getting feed by URL %s: %w", nextFeedURL, err)
	}

	fmt.Println("New Feed to Update: %s", nextFeedData.Name)
	err = s.db.MarkFeedAsFetched(context.Background(), nextFeedData.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	newFeedData, err := fetchFeed(context.Background(), nextFeedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	for _, item := range newFeedData.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
	}
	// data, err := json.MarshalIndent(newFeedData, "", "  ")
	// if err != nil {
	// 	log.Fatalf("error marshaling feed to JSON: %v", err)
	// }
	// fmt.Println(string(data))
	//
	return nil
}
