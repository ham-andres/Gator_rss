package main

import (
	"fmt"
	"log"
	"os"
	"github.com/ham-andres/Gator_rss/internal/config"

)



func main()  {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v",err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	myState := state{cfg: &cfg}

	cmds := commands{handlers: map[string]func(*state, command) error{}}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}
	exec := command{name: os.Args[1], arguments: os.Args[2:]}
	err = cmds.run(&myState, exec)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
