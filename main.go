package main

import (
	"fmt"

	"github.com/hatimhas/gator_rss/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("Error opening config in main(): %v\n", err)
		return
	}
	config.SetUser("hatim")

	fmt.Println(config.PrettyString())
}
