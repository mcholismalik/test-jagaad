# test-jagaad
Technical Test Jagaad
- Command : help, fetch, search, exit

## Pre Requisite
- Go version 1.20

## How To Run
``` bash
# run 
$ go run main.go

# help: Display this help message 
$ > help

# fetch: Fetch users and save to CSV 
$ > fetch

# search: Search users by tag from CSV 
$ > search --tags=tag1,tag2

# exit: Exit the CLI 
$ > exit
```

## Architecture 
This project built in clean architecture that contains some layer :
1. Entity   
2. Domain 
3. Usecase
4. Cmd

# Author
mcholismalik.official@gmail.com
