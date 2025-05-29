package main

import (
	"fmt"

	"github.com/hatimhas/gator_rss/internal/config"
	"github.com/hatimhas/gator_rss/internal/database"
)

// psql postgres://postgres:postgres@localhost:5432/gator
// goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
// sudo systemctl start postgresql
// sqlc generate
type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

// - This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	cmdFunc := c.cmdMap[cmd.name]
	if cmdFunc == nil {
		return fmt.Errorf("command %s not found", cmd.name)
	}
	// Passed cmd to cmdFunc to acess arguments
	return cmdFunc(s, cmd)
}

// - This method registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) error {
	if _, exists := c.cmdMap[name]; exists {
		return fmt.Errorf("error: command already exists")
	}
	c.cmdMap[name] = f
	return nil
}
