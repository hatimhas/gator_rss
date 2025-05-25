package main

import "github.com/hatimhas/gator_rss/internal/config"

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}
