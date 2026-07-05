package storage

import (
	"os"
	"passwordmanagercli/models"
)

func createPasswordFile() ([]models.Password, error) {
	err := os.WriteFile(passwordFilePath, []byte("[]"), 0644)
	if err != nil {
		return nil, err
	}
	return []models.Password{}, nil
}
