package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hatimhas/gator_rss/internal/database"
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

func scrapeFeed(s *state) error {
	nextFeedURL, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	nextFeedData, err := s.db.GetFeedByURL(context.Background(), nextFeedURL)
	if err != nil {
		return fmt.Errorf("error getting feed by URL %s: %w", nextFeedURL, err)
	}

	fmt.Printf("New Feed to Update: %s\n", nextFeedData.Name)
	err = s.db.MarkFeedAsFetched(context.Background(), nextFeedData.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	newFeedData, err := fetchFeed(context.Background(), nextFeedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	for _, item := range newFeedData.Channel.Item {
		var publishedAt time.Time
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = t
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			FeedID:      nextFeedData.ID,
			Title:       item.Title,
			Description: item.Description,
			Url:         item.Link,
			PublishedAt: publishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				fmt.Printf("Post %s already exists, skipping...\n", item.Title)
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
		fmt.Printf("Post %s created successfully\n", item.Title)
	}
	fmt.Printf("Feed %s collected, %v posts found\n", newFeedData.Channel.Title, len(newFeedData.Channel.Item))
	return nil
}
