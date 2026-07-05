package main

import (
	"fmt"
	"os"
	"passwordmanagercli/commands"
	"passwordmanagercli/util"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		util.PrintUsage()
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Please provide service, username and password")
			util.PrintUsage()
			return
		}

		service := os.Args[2]
		username := os.Args[3]
		password := os.Args[4]

		commands.AddPassword(service, username, password)

	case "list":
		commands.ListPasswords()

	case "search":
		if len(os.Args) != 3 {
			fmt.Println("Please provide service name")
			util.PrintUsage()
			return
		}

		service := os.Args[2]
		commands.SearchPassword(service)

	default:
		fmt.Println("Unknown command.")
		util.PrintUsage()
	}
}
