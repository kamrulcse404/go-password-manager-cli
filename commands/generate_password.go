package commands

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

func GeneratePassword(length int) {
	if length < 8 {
		fmt.Println("Password length must be at least 8.")
		return
	}

	password := make([]byte, length)
	charsetLen := len(charset)

	for idx := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(charsetLen)))
		if err != nil {
			fmt.Println("Random number generate problem:", err)
			return
		}
		password[idx] = charset[n.Int64()]
	}

	fmt.Printf("Your new generated password is: %s\n", string(password))
}
