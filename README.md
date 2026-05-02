- currently doing config file, will describe later.
  -- linking three dots in config:
  1. we wrote three functions
    -read() - loads the saved file into memory
    -SetUser() - changes something in the memory, change user is a method of a struct.
    -write() - it is a helper function for SetUser().

2. we created command.go file
  - created 3 struct type:- state, command and commands
  - inside config pointer, in command we wrote name, arguments slice[] and handlers map inside the commands to store the used commands.
  - handlerLogin - where we SetUser the username which is available in arguments
  
  - run function and register function - as the name suggested. 

3. inside the main - 
  - we firstly read config
  - create state and shared the address of &config
  - commands handlers ( where we register login command)
  - and lastly execute the command


-  sql part:

1. connection string for postgres: 
"postgres://postgres:postgres@localhost:5432/gator"

- interaction with postgres through client like psql 
  usage (start): sudo service postgresql start
  installation: sudo apt install postgresql postgresql-contrib

  start: sudo service postgresql start
  stop: sudo service postgresql stop
  connect: sudo -u postgres psql

  - Goose migration 
  Just a .sql file with some SQL queries and some special comments.
  CREATE: goose <connection string> up 
  DROP: goose <connection string> down 

  - SQLC 
  Go program that generates Go code from SQL queries like ORM but not exactly ORM.

  2. step to handle the table and their entries 
  - added Reset functionality (handler_reset.go) 
  all the function have similar signature func <name> (parameter 1(state), parameter2 (command)) error
  - added handler_user.go file which has 3 function to manage user's info, login , register, listing users (showing current user name) 
  
 3. RSS part
 - used a rss struct which is output for the fetchFeed function
 used: http.NewRequestWithContext
       req.Header.Set("User-Agent","gator")
       made client:= &http.Client{..}
       used resp,err := Client.do(req)
       io.ReadAll(resp.Body)
       feed := New(RSSFeed)
       err = xml.Unmarshal

       this above part was similar to http 

       html.UnescapeString to make it human readable 
4. Under RSS part (Feed)
- handlerFeed func and handlerAgg function
- handlerAgg: fetchFeed(ctx, Url )
- handlerFeed: uses s.db.GetUser(ctx, CreateFeedParams struct)


