package main
import (
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/ham-andres/Gator_rss/internal/config"
	"github.com/ham-andres/Gator_rss/internal/database"
	_ "github.com/lib/pq"

)



func main()  {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v",err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	
	// loading database
	db, err := sql.Open("postgres",cfg.DURL)
	if err != nil {
		log.Fatalf("error loading database: %v",err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	myState := state{db: dbQueries, cfg: &cfg}


	cmds := commands{handlers: map[string]func(*state, command) error{}}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
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
