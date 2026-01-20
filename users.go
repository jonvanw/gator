package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}
	fmt.Println("Users:")
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf(" *  %s (current)\n", user.Name)
		} else {
			fmt.Printf(" *  %s\n", user.Name)
		}
	}
	return nil
}