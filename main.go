package main

import (
	"fmt"
	"log"
	
	"github.com/ham-andres/Gator_rss/internal/config"

)



func main()  {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	cfg.SetUser("hamandres")
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n",cfg)
}
