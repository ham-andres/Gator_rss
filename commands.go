package main

import (
	"github.com/ham-andres/Gator_rss/internal/config"
	"fmt"
)

type state struct {
	
	cfg *config.Config
}

type command struct {
	name	string
	arguments	[]string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no username given")
	}
	
	name := cmd.arguments[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	
	fmt.Printf("%v user has been set\n", name)
	return nil

}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if callFunc, exists := c.handlers[cmd.name]; exists {
		return callFunc(s, cmd)
		
	}
	return fmt.Errorf("command is not registered")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
