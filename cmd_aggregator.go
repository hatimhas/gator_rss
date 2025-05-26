package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func handlerAggregator(s *state, cmd command) error {
	feedData, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	// Print the feed title and description

	data, err := json.MarshalIndent(feedData, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling feed to JSON: %v", err)
	}
	fmt.Println(string(data))
	return nil
}
