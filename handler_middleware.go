package main

import (
	"context"

	"github.com/ham-andres/Gator_rss/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state,cmd command) error {
		currUserName := s.cfg.CurrentUserName
		user, err := s.db.GetUser(context.Background(), currUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
