package util

import (
	"passwordmanagercli/models"
)

func GetNextID(passwords []models.Password) int {
	highestID := 0

	for _, password := range passwords {
		if password.ID >  highestID {
			highestID = password.ID
		}
	}
	return  highestID + 1
}