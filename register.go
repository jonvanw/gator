package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jonvanw/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("register command requires exactly one argument, the username.")
	}

	userName := cmd.args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: userName,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	fmt.Printf("User created successfully:\n")
	fmt.Printf("  ID: %s\n", user.ID)
	fmt.Printf("  Name: %s\n", user.Name)
	fmt.Printf("  CreatedAt: %s\n", user.CreatedAt.Format(time.RFC3339))
	fmt.Printf("  UpdatedAt: %s\n", user.UpdatedAt.Format(time.RFC3339))

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("failed to login as new user: %v", err)
	}

	fmt.Printf("Logged in as new user %s.\n", userName)

	return nil
}