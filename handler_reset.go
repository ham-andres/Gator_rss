package main

import (
	"context"
	"fmt"	
)

func handlerReset(s *state, cmd command) error {
		err := s.db.DeleteAllUsers(context.Background())
		if err != nil {
			return fmt.Errorf("reset database unsuccessful: %w", err)
		}
		fmt.Println("Reset successful")
		return nil
}
