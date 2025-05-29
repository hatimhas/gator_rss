package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hatimhas/gator_rss/internal/database"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("follow cmd requires one arg: feed_url")
	}
	feedURL := cmd.arguments[0]

	// use query to get feed_id using given feed_url, then use current logged in user_id, use CreateFeedFollow query to create a new feed follow record

	feed_data, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error getting feed by URL %s: %v", feedURL, err)
	}

	cur_user_data, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user %s: %v", s.config.CurrentUserName, err)
	}

	feeds_users_data, err := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    cur_user_data.ID,
			FeedID:    feed_data.ID,
		})
	if err != nil {
		return fmt.Errorf("error following feed %s: %v", feedURL, err)
	}

	fmt.Printf("Successfully followed FeedName: %s with User: %s \n", feeds_users_data.FeedName, feeds_users_data.UserName)
	return nil
}
