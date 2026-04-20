package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rssFeed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil
}
