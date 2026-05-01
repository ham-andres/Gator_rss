package main 

import (
	"fmt"
	"context"

)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Fetching failed: %v",err)
	}

	fmt.Printf("%+v\n",feed)
	return nil
}
