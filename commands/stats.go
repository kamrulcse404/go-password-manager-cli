package commands

import (
	"fmt"
	"passwordmanagercli/storage"
	"strings"
)

func Stats() {
	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("Error loading passwords!", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("No passwords found.")
		return
	}

	totalPasswords := len(passwords)
	longestPassword := 0
	shortestPassword := len(passwords[0].Password)
	services := make(map[string]bool)
	serviceCount := make(map[string]int)

	for _, password := range passwords {
		services[strings.ToLower(password.Service)] = true
		serviceCount[strings.ToLower(password.Service)]++
		if len(password.Password) > longestPassword {
			longestPassword = len(password.Password)
		}

		if len(password.Password) < shortestPassword {
			shortestPassword = len(password.Password)
		}
	}
	uniqueServices := len(services)

	mostUsedService := ""
	mostUsedCount := 0

	for service, count := range serviceCount {
		if count > mostUsedCount {
			mostUsedCount = count
			mostUsedService = service
		}
	}

	fmt.Println("Password Manager Statistics")
	fmt.Println("-----------------------------------")
	fmt.Printf("Total Passwords        : %d\n", totalPasswords)
	fmt.Printf("Unique Services        : %d\n", uniqueServices)
	fmt.Printf("Longest Password       : %d characters\n", longestPassword)
	fmt.Printf("Shortest Password      : %d characters\n", shortestPassword)
	fmt.Printf("Most Used Service      : %s (%d accounts)\n", mostUsedService, mostUsedCount)
}
