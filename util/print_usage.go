package util

import "fmt"

func PrintUsage() {
	fmt.Println("Password Manager CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  add <service> <username> <password>   Add a new password")
}