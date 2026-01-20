package main

import (
	"github.com/jonvanw/gator/internal/config"
)

type state struct {
	config *config.Config
}

func NewState(cfg *config.Config) *state {
	return &state{
		config: cfg,
	}
}
