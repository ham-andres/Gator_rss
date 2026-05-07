package main 

import (
	 "context"
   "fmt"
   "time"

   "github.com/ham-andres/Gator_rss/internal/database"
   "github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %v <url>",cmd.name)
	}
	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed accessing the feed: %w", err)
	}

	// user code deleted, use middlewareLoggedIn

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

func handlerFollowing(s *state, cmd command, user database.User) error {
// user code deleted, uses middlewareLoggedIn	
	currUserId := user.ID 
	
	followingFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currUserId) 
	if err != nil {
		return fmt.Errorf("failed accessing feed for current user: %w",err)
	}
	if len(followingFeeds) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	for _,fFeed := range followingFeeds {
		fmt.Printf("Feed Name: %v \n", fFeed.FeedName)
		fmt.Println("----------------------------")
	}
	return nil
}


// unfollow command
func handlerUnfollow(s *state, cmd command, user database.User) error {
	url := cmd.arguments[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed accessing feed to unfollow: %w", err)
	}

	unfollowParams := database.UnfollowFeedParams {
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.UnfollowFeed(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("failed unfollow from database: %w",err)
	}
	return nil
}


