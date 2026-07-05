package commands

import (
	"fmt"
	"passwordmanagercli/models"
	"passwordmanagercli/storage"
	"passwordmanagercli/util"
)

func AddPassword(service, userNameOrEmail, password string) {
	if service == "" || userNameOrEmail == "" || password == "" {
		fmt.Println("Password service, username/email and password cannot be empty.")
		return
	}

	passwords, err := storage.LoadPasswords()

	if err != nil {
		fmt.Println("Error loading passwords!", err)
		return
	}

	newPassword := models.Password{
		ID:       util.GetNextID(passwords),
		Service:  service,
		Username: userNameOrEmail,
		Password: password,
	}

	passwords = append(passwords, newPassword)
	err = storage.SavePasswords(passwords)

	if err != nil {
		fmt.Println("Error adding passwords! ", err)
		return
	}

	fmt.Printf("Password #%d added successfully.\n", newPassword.ID)
}
