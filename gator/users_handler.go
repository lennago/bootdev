package main

import (
	"context"
	"fmt"

	"github.com/lennago/bootdotdev_gator/internal/database"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error listing users: %w", err)
	}
	printUsers(users, s.cfg.CurrentUserName)
	return nil
}

func printUsers(users []database.User, currentUser string) {
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
}
