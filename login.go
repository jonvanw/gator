package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("login command requires exactly one argument, the username.")
	}

	userName := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user %s does not exist, please register first", userName)
		}
		return fmt.Errorf("failed to get user: %v", err)
	}

	s.cfg.SetUser(userName)

	fmt.Printf("Logged in as %s\n", userName)

	return nil
}