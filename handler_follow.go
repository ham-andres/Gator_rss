package main 

import (
	 "context"
   "fmt"
   "time"

   "github.com/ham-andres/Gator_rss/internal/database"
   "github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <url>",cmd.name)
	}
	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed accessing the feed: %w", err)
	}
	currName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currName)
	if err != nil {
		return fmt.Errorf("failed accessing current username: %w",err)
	}
	timestamp := time.Now()
	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: 	timestamp,
		UpdatedAt: 	timestamp,
		UserID: 	user.ID,
		FeedID: 	feed.ID,
	}

	feedRow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("Feed Name: %v \n", feedRow.FeedName)
	fmt.Printf("User Name: %v \n", feedRow.UserName)
	fmt.Println("=========================")
	return nil

}

func handlerFollowing(s *state, cmd command) error {
	currUserName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currUserName)
	if err != nil {
		return fmt.Errorf("failed accessing user id: %w", err)
	}
	currUserId := user.ID 
	
	followingFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currUserId) 
	if err != nil {
		return fmt.Errorf("failed accessing feed for current user: %w",err)
	}
	for _,fFeed := range followingFeeds {
		fmt.Printf("Feed Name: %v \n", fFeed.FeedName)
		fmt.Println("----------------------------")
	}
	return nil
}
