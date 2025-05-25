package main

import (
	"fmt"
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
	if err := s.config.SetUser(username); err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}
	// Print a message to the terminal that the user has been set.
	fmt.Printf("User set to %s\n", username)
	return nil
}
