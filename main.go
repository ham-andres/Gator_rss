package main

import (
	"fmt"
	"log"
	
	"github.com/ham-andres/Gator_rss/internal/config"

)



func main()  {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v",err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("hamandres")
	if err != nil {
		log.Fatalf("coundn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("err reading config: %v",err)
	}
	fmt.Printf("Read config again: %+v\n",cfg)
}
