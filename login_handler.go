package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	if _, err := s.db.GetUser(context.Background(), cmd.Args[0]); err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}
