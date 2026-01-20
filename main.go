package main

import (
	"fmt"
	"log"

	"github.com/jonvanw/gator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	cfg.SetUser("Jon")

	cfg, err = config.ReadConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	fmt.Printf("Contents of config file: %+v\n", cfg)
}