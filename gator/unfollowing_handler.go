package main

import (
	"context"
	"fmt"

	"github.com/lennago/bootdotdev_gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, curUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	if err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: curUser.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}
	fmt.Printf("%s unfollowed successfully!\n", feed.Name)
	return nil
}
