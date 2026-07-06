package commands

import (
	"fmt"
	"passwordmanagercli/crypto"
	"passwordmanagercli/models"
	"passwordmanagercli/storage"
	"passwordmanagercli/util"
	"strings"
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

	for _, password := range passwords {
		if (strings.ToLower(password.Service) == strings.ToLower(service)) && (strings.ToLower(password.Username) == strings.ToLower(userNameOrEmail)) {
			fmt.Println("Service already exists with same username!")
			return
		}
	}

	encryptedPassword, err := crypto.Encrypt(password)
	if err != nil {
		fmt.Println("Failed to encrypt password:", err)
		return
	}

	newPassword := models.Password{
		ID:       util.GetNextID(passwords),
		Service:  service,
		Username: userNameOrEmail,
		Password: encryptedPassword,
	}

	passwords = append(passwords, newPassword)
	err = storage.SavePasswords(passwords)

	if err != nil {
		fmt.Println("Error adding passwords! ", err)
		return
	}

	fmt.Printf("Password #%d added successfully.\n", newPassword.ID)
}
