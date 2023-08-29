package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/test-jagaad/internal/usecase"
)

type (
	Command interface {
		Init()
	}

	command struct {
		userUc usecase.UserUc
	}
)

func NewCommandCmd(userUc usecase.UserUc) Command {
	return &command{userUc}
}

func (c *command) Init() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("CLI Test Jagaad")
	fmt.Println("Type 'help' for available commands")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Split(input, " ")
		command := args[0]
		switch command {
		case "help":
			c.help()
		case "fetch":
			c.fetch()
		case "search":
			c.search(args)
		case "exit":
			fmt.Println("Exiting CLI...")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

func (c *command) help() {
	fmt.Println("Available commands:")

	commands := []string{
		"help: Display this help message",
		"fetch: Fetch users and save to CSV",
		"search: Search users by tag from CSV",
		"exit: Exit the CLI",
	}
	for _, cmd := range commands {
		fmt.Println(cmd)
	}
}

func (c *command) fetch() {
	fmt.Println("Fetching users and save to csv...")

	resp, err := c.userUc.Fetch()
	if err != nil {
		fmt.Println("Error fetch users and save to csv")
	} else {
		for _, data := range resp.Details {
			status := "success"
			if data.Err != nil {
				status = "error"
			}

			msg := fmt.Sprintf(`Endpoint %s: %s`, data.Name, status)
			fmt.Println(msg)
		}
		msg := fmt.Sprintf(`Filename: %s`, resp.Filename)
		fmt.Println(msg)
	}
}

func (c *command) search(args []string) {
	tags := []string{}
	if len(args) < 2 || !strings.HasPrefix(args[1], "--tags=") {
		fmt.Println("Usage: search --tags=mock1,mock2")
	} else {
		tagStr := strings.TrimPrefix(args[1], "--tags=")
		tags = strings.Split(tagStr, ",")
		fmt.Println("Searching users with tags:", tags)
	}

	resp, err := c.userUc.Search(tags)
	if err != nil {
		fmt.Println("Error searching users with tags:", tags)
	} else {
		fmt.Println("Found users", resp)
	}
}
