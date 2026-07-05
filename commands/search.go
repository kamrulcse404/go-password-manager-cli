package commands

import (
	"fmt"
	"passwordmanagercli/storage"
	"strings"
)

func SearchPassword(service string) {
	if service == "" {
		fmt.Println("service cannot be empty")
		return
	}

	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("failed to load passwords: ", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("Passwords not found.")
		return
	}

	found := false

	for _, password := range passwords {
		if strings.EqualFold(password.Service, service) {
			found = true
			fmt.Printf("ID      : %d\n", password.ID)
			fmt.Printf("Service : %s\n", password.Service)
			fmt.Printf("Username: %s\n", password.Username)
			fmt.Printf("Password: %s\n", password.Password)
		}
	}

	if !found {
		fmt.Println("Service not found.")
		return
	}
}
