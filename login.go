package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("login command requires exactly one argument, the username.")
	}

	userName := cmd.args[0]

	s.config.SetUser(userName)

	fmt.Printf("Logged in as %s\n", userName)

	return nil
}