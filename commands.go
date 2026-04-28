package main

import (
	"time"
	"context"
	"fmt"

	"github.com/google/uuid"
  "os"
	"github.com/ham-andres/Gator_rss/internal/config"
	"github.com/ham-andres/Gator_rss/internal/database"
)

type state struct {
	db	*database.Queries
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
	ctx := context.Background()
	name := cmd.arguments[0]
	user, err := s.db.GetUser(ctx, name)
	if err != nil {
		fmt.Printf("error fetching user in login: %v\n", err)
		os.Exit(1)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %v\n",err) 
	}
	
	fmt.Printf("%v user has been set\n",user.Name)
	return nil

}



func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no username given")
	}
	ctx := context.Background()
	name := cmd.arguments[0]
	userParam := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}
	user, err := s.db.CreateUser(ctx, userParam)
	if err != nil {
		fmt.Printf("error creating user: %v\n",err)
		os.Exit(1)
	}


	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("%v user has been created\n",user.Name)
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
