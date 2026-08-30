package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if f, ok := c.registeredCommands[cmd.Name]; !ok {
		return errors.New("command not found")
	} else {
		return f(s, cmd)
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
