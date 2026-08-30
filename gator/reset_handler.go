package main

import (
	"fmt"
	"os/exec"
)

const dir = "./sql/schema"

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 && !(len(cmd.Args) == 1 && cmd.Args[0] == "up") {
		return fmt.Errorf("usage: %s [up]", cmd.Name)
	}
	if len(cmd.Args) == 0 {
		migrate_down := exec.Command("goose", "postgres", s.cfg.DBURL, "reset")
		migrate_down.Dir = dir
		if _, err := migrate_down.Output(); err != nil {
			return err
		}
	}
	migrate_up := exec.Command("goose", "postgres", s.cfg.DBURL, "up")
	migrate_up.Dir = dir
	if _, err := migrate_up.Output(); err != nil {
		return err
	}
	if len(cmd.Args) == 0 {
		fmt.Println("Database reset successfully!")
	} else {
		fmt.Println("Database upgraded successfully!")
	}
	return nil
}
