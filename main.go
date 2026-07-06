package main

import (
	"fmt"
	"os"
	"passwordmanagercli/commands"
	"passwordmanagercli/util"
	"strconv"
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

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Please provide ID")
			util.PrintUsage()
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID. Please enter a valid number.")
			return
		}

		commands.RemovePassword(id)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Please provide ID & new password")
			util.PrintUsage()
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID. Please enter a valid number.")
			return
		}

		newPassword := os.Args[3]

		commands.UpdatePassword(id, newPassword)

	case "stats":
		commands.Stats()

	case "generate":
		if len(os.Args) < 3 {
			fmt.Println("Please provide password length")
			util.PrintUsage()
			return
		}

		len, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid number. Please enter a valid number [range(1-70)]")
			return
		}

		commands.GeneratePassword(len)

	default:
		fmt.Println("Unknown command.")
		util.PrintUsage()
	}
}
