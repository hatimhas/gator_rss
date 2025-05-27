package main

import (
	"context"
	"fmt"
)

func handlerPrintAllFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("feeds cmd requires no arg")
	}

	feedsData, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting all feeds data: %v", err)
	}

	for i := range feedsData {
		fmt.Printf("* Feed Name:          %s\n", feedsData[i].FeedName)
		fmt.Printf("* Feed URL:           %s\n", feedsData[i].FeedUrl)
		fmt.Printf("* Feed Creator:        %s\n", feedsData[i].FeedCreatorName)
	}

	return nil
}
