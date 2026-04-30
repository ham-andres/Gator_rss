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



