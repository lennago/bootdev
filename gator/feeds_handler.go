package main

import (
	"context"
	"fmt"

	"github.com/lennago/bootdotdev_gator/internal/database"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error listing feeds: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	fmt.Printf("Found %d feeds:\n", len(feeds))
	if err := printFeeds(feeds, s.db); err != nil {
		return fmt.Errorf("error printing feeds: %w", err)
	}
	return nil
}

func printFeeds(feeds []database.Feed, db *database.Queries) error {
	for _, feed := range feeds {
		user, err := db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeed(feed, user)
		println("=====================================")
	}
	return nil
}
