package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	pokemon, ok := cfg.caught[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
