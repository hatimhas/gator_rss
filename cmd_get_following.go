package main

import (
	"context"
	"fmt"
)

func handlerGetFollowing(s *state, cmd command) error {
	currentUser := s.config.CurrentUserName

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
