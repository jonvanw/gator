package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to clear users: %v", err)
	}
	fmt.Println("Users cleared successfully.")
	return nil
}