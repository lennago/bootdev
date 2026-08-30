package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.caught {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
