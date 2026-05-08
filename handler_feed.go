package main

import (
    "context"
    "fmt"
    "time"

    "github.com/ham-andres/Gator_rss/internal/database"
    "github.com/google/uuid"
)

func handlerFeed(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) != 2 {
		return fmt.Errorf("usage: %v <name> <Url> ", cmd.name)
	}	
// user code deleted uses middle ware func

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
	currTime := time.Now()
	addFeedFollowParams := database.CreateFeedFollowParams{
		ID: 	uuid.New(),
		CreatedAt: 	currTime,
		UpdatedAt: 	currTime,
		UserID: 	user.ID,
		FeedID: 	feed.ID,
	}

	addFeedFollow, err := s.db.CreateFeedFollow(context.Background(), addFeedFollowParams)
	if err != nil {
		return fmt.Errorf("failed adding feed of user: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Printf("Feed added: %v for user: %v \n",addFeedFollow.FeedName, addFeedFollow.UserName)
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
	fmt.Printf("* LastFetchedAt: 	 %v\n",feed.LastFetchedAt.Time)
}

func handlerShowFeeds(s *state, cmd command) error {
	feed, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _,f := range feed {
		user, err := s.db.GetUserById(context.Background(),f.UserID.UUID)
		if err != nil {
			return fmt.Errorf("couldn't access username %w", err)
		}
		fmt.Printf("Username: %v", user.Name)
		printFeed(f)
		fmt.Println("--------------------------")
	}
	fmt.Println("=====================================")
	return nil
}
