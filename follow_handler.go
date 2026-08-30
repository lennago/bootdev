package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lennago/bootdotdev_gator/internal/database"
)

func handlerFollow(s *state, cmd command, curUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	feedFollow, err := s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    curUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed follow created:")
	printFeedFollows(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func printFeedFollows(username string, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
