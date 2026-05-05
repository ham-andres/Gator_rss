package main

import (
    "context"
    "fmt"
    "time"

    "github.com/ham-andres/Gator_rss/internal/database"
    "github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {

	if len(cmd.arguments) != 2 {
		return fmt.Errorf("usage: %v <name> <Url> ", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), 
			database.CreateFeedParams{
				ID:	uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Name: cmd.arguments[0],
				Url: cmd.arguments[1],
				UserID: uuid.NullUUID{UUID: user.ID, Valid: true}, // this cause me hell of a problem figuring out
			})
	if err != nil {
		return fmt.Errorf("Error while accessing feed: %w", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
