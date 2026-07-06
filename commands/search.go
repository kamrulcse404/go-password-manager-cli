package commands

import (
	"fmt"
	"passwordmanagercli/storage"
	"strings"
	"passwordmanagercli/crypto"
)

func SearchPassword(service string) {
	if service == "" {
		fmt.Println("service cannot be empty")
		return
	}

	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("Failed to load passwords: ", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("No passwords found.")
		return
	}

	found := false
	fmt.Printf("%-5s %-15s %-20s %-30s\n", "ID", "Service", "Username", "Password")
	fmt.Println("------------------------------------------------------")

	for _, password := range passwords {
		if strings.EqualFold(password.Service, service) {
			decryptedText, err := crypto.Decrypt(password.Password)
			if err != nil {
				fmt.Println("Failed to decrypt password:", err)
				return
			}
			found = true
			fmt.Printf("%-5d %-15s %-20s %-30s\n",
				password.ID,
				password.Service,
				password.Username,
				decryptedText,
			)
		}
	}

	if !found {
		fmt.Println("Service not found.")
		return
	}
}
