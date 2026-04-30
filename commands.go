package main

import (
	"errors"
)



type command struct {
	name	string
	arguments	[]string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}




func (c *commands) run(s *state, cmd command) error {
	if callFunc, exists := c.handlers[cmd.name]; exists {
		return callFunc(s, cmd)
		
	}
	return errors.New("command is not registered")
}


