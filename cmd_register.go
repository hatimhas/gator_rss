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

func handlerRegister(s *state, cmd command) error {
	// If the command's arg's slice is empty, return an error; the login handler expects a single argument, the username.
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("regiser command requires a name argument")
	}
	username := strings.TrimSpace(cmd.arguments[0])
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}
	getUser, err := s.db.GetUser(context.Background(), userParams.Name)
	if err == nil {
		if getUser != (database.User{}) {
			fmt.Printf("User %s already exists with ID %s\n", getUser.Name, getUser.ID)
			os.Exit(1)
		}
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	s.config.SetUser(user.Name)
	fmt.Printf("User %s created with ID %s\n", user.Name, user.ID)
	return nil
}
