package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(st *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		fmt.Println("unknown command:", cmd.Name)
	}

	return f(st, cmd)
}
