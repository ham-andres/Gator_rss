package main 

import (
	"fmt"
	"context"

)

func handleAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Fetching failed: %v",err)
	}

	fmt.Printf("%+v",feed)
	return nil
}
