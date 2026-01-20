package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jonvanw/gator/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config: %v\n", err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	s := NewState(cfg)

	commands := NewCommands()
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{name: cmdName, args: cmdArgs}

	err = commands.run(s, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}