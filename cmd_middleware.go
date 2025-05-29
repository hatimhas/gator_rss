package main

import (
	"context"
	"fmt"

	"github.com/hatimhas/gator_rss/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Fetch user from state/db here
		// If error, return error immediately
		// Otherwise, call: handler(s, cmd, user)
		userData, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting current user data: %v", err)
		}
		return handler(s, cmd, userData)
	}
}
