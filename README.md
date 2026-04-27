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

