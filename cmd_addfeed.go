package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hatimhas/gator_rss/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("addfeed command requires a name and url argument")
	}
	feedName := strings.TrimSpace(cmd.arguments[0])
	feedURL := strings.TrimSpace(cmd.arguments[1])

	currentUser := s.config.CurrentUserName
	currentUserData, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error getting data of current logged from users db: %v", err)
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       feedURL,
		Name:      feedName,
		UserID:    currentUserData.ID,
	}

	getFeed, err := s.db.GetFeedByURL(context.Background(), feedParams.Url)
	if err == nil {
		if getFeed != (database.Feed{}) {
			fmt.Printf("Feed %s already exists with ID %s\n", feedName, getFeed.ID)
			os.Exit(1)
		}
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating new feed: %w", err)
	}
	fmt.Printf("Feed %s created with URL %s\n", feed.Name, feed.Url)
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
