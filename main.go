package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hatimhas/gator_rss/internal/config"
	"github.com/hatimhas/gator_rss/internal/database"
	_ "github.com/lib/pq"
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
		log.Fatalf("error reading config : %v", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := commands{cmdMap: make(map[string]func(*state, command) error)}
	if err := cmds.register("login", handlerLogin); err != nil {
		fmt.Println(err)
	}

	if err := cmds.register("register", handlerRegister); err != nil {
		fmt.Println(err)
	}
	if err := cmds.register("reset", handlerReset); err != nil {
		fmt.Println(err)
	}
	if err := cmds.register("users", handlerGetUsers); err != nil {
		fmt.Println(err)
	}
	if err := cmds.register("agg", handlerAggregator); err != nil {
		fmt.Println(err)
	}
	if err := cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed)); err != nil {
		fmt.Println(err)
	}
	if err := cmds.register("feeds", handlerPrintAllFeeds); err != nil {
		fmt.Println(err)
	}
	if err := cmds.register("following", middlewareLoggedIn(handlerGetFollowing)); err != nil {
		fmt.Println(err)
	}
	if err := cmds.register("follow", middlewareLoggedIn(handlerFollowFeed)); err != nil {
		fmt.Println(err)
	}
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Fprintf(os.Stderr, "Error running command %s: %v\n", cmd.name, err)
		os.Exit(1)
	}
}
