package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		log.Fatalf("couldn't delete users: %v", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
