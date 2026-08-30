package main

import (
	"context"
	"fmt"

	"github.com/lennago/bootdotdev_gator/internal/database"
)

func handlerFollowing(s *state, cmd command, curUser database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), curUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}
	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}
	fmt.Printf("Feed follows for user %s:\n", curUser.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}
	return nil
}
