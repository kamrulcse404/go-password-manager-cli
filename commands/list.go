package commands

import (
	"fmt"
	"passwordmanagercli/crypto"
	"passwordmanagercli/storage"
)

func ListPasswords() {
	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("Error loading passwords: ", err)
		return
	}

	if len(passwords) == 0 {
		fmt.Println("No passwords found")
		return
	}

	fmt.Printf("%-5s %-15s %-20s %-30s\n", "ID", "Service", "Username", "Password")
	fmt.Println("-------------------------------------------------------------------")

	for _, password := range passwords {
		decryptedText, err := crypto.Decrypt(password.Password)
		if err != nil {
			fmt.Println("Failed to decrypt password:", err)
			return
		}
		fmt.Printf("%-5d %-15s %-20s %-30s\n",
			password.ID,
			password.Service,
			password.Username,
			decryptedText,
		)
	}
}
