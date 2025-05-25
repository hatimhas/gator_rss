package main

import (
	"fmt"

	"github.com/hatimhas/gator_rss/internal/config"
)

//
// func main() {
// 	config, err := config.Read()
// 	if err != nil {
// 		fmt.Printf("Error opening config in main(): %v\n", err)
// 		return
// 	}
// 	config.SetUser("hatim")
//
// 	fmt.Println(config.PrettyString())
// }
//
//

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config at scarth.go: %v", err)
		return
	}
	s := state{config: &cfg}
	cmd := command{name: "login", arguments: []string{"harim"}}

	if err := handlerLogin(&s, cmd); err != nil {
		fmt.Printf("error handling Login at scarth.go: %v", err)
		return

	}
}
