package main

import (
	"github.com/jonvanw/gator/internal/config"
	"github.com/jonvanw/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func NewState(cfg *config.Config, db *database.Queries) *state {
	return &state{
		cfg: cfg,
		db: db,
	}
}
