package main

import (
	"context"
	"fmt"
	"time"  // for UserParam

	"github.com/ham-andres/Gator_rss/internal/database"
	"github.com/google/uuid"
)

// Register function 

func handlerRegister(s *state, cmd command) error {

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}

	name := cmd.arguments[0]
	userParam := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}


	user, err := s.db.CreateUser(context.Background(), userParam)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: %v <name>", cmd.name)
	}
	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	user, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("table is empty: %w", err)
	}

	for _, u := range(user) {
		if s.cfg.CurrentUserName == u.Name {
			fmt.Printf("* %v (current)\n",u.Name)
		} else {
			fmt.Printf("* %v\n",u.Name)
		}
	}
	return nil
}



func printUser( user database.User) {
	fmt.Printf(" * ID:	%v\n", user.ID)
	fmt.Printf(" * Name:	%v\n", user.Name)
}
