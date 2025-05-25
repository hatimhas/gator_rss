package main

import (
	"fmt"
	"os"

	"github.com/hatimhas/gator_rss/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: not enough arguments provided.")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{name: cmdName, arguments: cmdArgs}

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config : %v", err)
		return
	}

	s := state{config: &cfg}

	cmds := commands{cmdMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if err := cmds.run(&s, cmd); err != nil {
		fmt.Fprintf(os.Stderr, "Error running command %s: %v\n", cmd.name, err)
		os.Exit(1)
	}
}
