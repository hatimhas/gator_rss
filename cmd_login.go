package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
)

func handlerLogin(s *state, cmd command) error {
	// If the command's arg's slice is empty, return an error; the login handler expects a single argument, the username.
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("login command requires a username argument")
	}
	// Use the state's access to the config struct to set the user to the given username. Remember to return any errors.
	username := strings.TrimSpace(cmd.arguments[0])
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("User %s does not exist\n", username)
			os.Exit(1)
		} else {
			return fmt.Errorf("error fetching user: %w", err)
		}
	}
	if err := s.config.SetUser(username); err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}
	// Print a message to the terminal that the user has been set.
	fmt.Printf("User set to %s\n", username)
	return nil
}
