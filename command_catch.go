package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Pokemon.Name)
	caught := rand.Intn(256) <= pokemon.Species.CaptureRate
	if caught {
		fmt.Printf("%s was caught!\n", pokemon.Pokemon.Name)
		cfg.caught[pokemon.Pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Pokemon.Name)
	}
	return nil
}
