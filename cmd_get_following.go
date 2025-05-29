package main

import (
	"context"
	"fmt"

	"github.com/hatimhas/gator_rss/internal/database"
)

func handlerGetFollowing(s *state, cmd command, user database.User) error {
	currentUser := user.Name

	followingData, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error getting following data for currentUser: %w", err)
	}

	fmt.Printf("Following for user %s:\n", currentUser)
	for i := range followingData {
		following := followingData[i]
		fmt.Printf("  %s \n", following.FeedName)
	}
	return nil
}
