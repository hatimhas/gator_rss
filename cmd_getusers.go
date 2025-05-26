package main

import (
	"context"
	"fmt"
	"os"
)

func handlerGetUsers(s *state, cmd command) error {
	users_data, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error fetching all users: %v\n", err)
		os.Exit(1)
	}
	for _, user := range users_data {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n",
				user.Name)
		}
	}
	return nil
}
