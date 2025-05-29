package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/hatimhas/gator_rss/internal/database"
)

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("unfollow command requires a feed URL argument")
	}
	currentUserID := user.ID
	feedURL := strings.TrimSpace(cmd.arguments[0])
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: currentUserID,
		Url:    feedURL,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed %s: %v", feedURL, err)
	}
	return nil
}
