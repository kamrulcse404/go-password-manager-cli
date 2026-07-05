package commands

import (
	"fmt"
	"passwordmanagercli/storage"
)

func ListPasswords() {
	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("Errro loading passwords: ", err)
		return 
	}

	if len(passwords) == 0 {
		fmt.Println("No passwords found")
		return
	}

	fmt.Printf("%-5s %-15s %-20s\n", "ID", "Service", "Username")
	fmt.Println("------------------------------------------")

	for _, password := range passwords {
		fmt.Printf("%-5d %-15s %-20s\n",
			password.ID,
			password.Service,
			password.Username,
		)
	}
}