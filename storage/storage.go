package storage

import (
	"encoding/json"
	"os"
	"passwordmanagercli/models"
)

const (
	dataDir          = "data"
	passwordFilePath = dataDir + "/passwords.json"
)

func LoadPasswords() ([]models.Password, error) {
	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return nil, err
		}
		return createPasswordFile()
	}

	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(passwordFilePath)

	if os.IsNotExist(err) {
		return createPasswordFile()
	}

	if err != nil {
		return nil, err
	}

	var passwords []models.Password

	err = json.Unmarshal(data, &passwords)
	if err != nil {
		return nil, err
	}

	return passwords, nil
}

func SavePasswords(passwords []models.Password) error {
	data, err := json.MarshalIndent(passwords, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(passwordFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
