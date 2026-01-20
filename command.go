package main

import (
	"fmt"
	"strings"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(s *state, cmd command) error
}

func NewCommands() *commands {
	return &commands{
		handlers: make(map[string]func(s *state, cmd command) error),
	}
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[strings.ToLower(cmd.name)]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(s *state, cmd command) error) {
	c.handlers[strings.ToLower(name)] = handler
}
