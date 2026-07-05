package commands

import (
	"fmt"
	"passwordmanagercli/storage"
)

func RemovePassword(id int) {
	if id < 1 {
		fmt.Println("Please provide a valid ID")
		return
	}

	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("Failed to load passwords: ", err)
		return
	}

	for idx, password := range passwords {
		if password.ID == id {
			passwords = append(passwords[:idx], passwords[idx+1:]...)

			err := storage.SavePasswords(passwords)
			if err != nil {
				fmt.Println("Error saving passwords:", err)
				return
			}
			fmt.Printf("Password #%d removed successfully.\n", id)
			return
		}
	}
	fmt.Printf("Password #%d not found.\n", id)
}
