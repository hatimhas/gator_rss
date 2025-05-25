package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Printf("Error deleting users: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All users have been deleted.")
	os.Exit(0)
	return nil
}
