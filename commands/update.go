package commands

import (
	"fmt"
	"passwordmanagercli/storage"
)

func UpdatePassword(id int, newPassword string) {
	if id < 1 {
		fmt.Println("Please provide a valid ID")
		return
	}

	if newPassword == "" {
		fmt.Println("Password cannot be empty.")
		return
	}

	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("Failed to load passwords: ", err)
		return
	}

	for idx := range passwords {
		if  passwords[idx].ID == id {
			passwords[idx].Password = newPassword

			err = storage.SavePasswords(passwords)

			if err != nil {
				fmt.Println("Error saving passwords:", err)
				return
			}

			fmt.Printf("Password #%d updated successfully!\n", id)
			return
		}
	}

	fmt.Printf("Password #%d not found.\n", id)
}
