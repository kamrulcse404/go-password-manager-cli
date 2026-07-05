package util

import "fmt"

func PrintUsage() {
	fmt.Println("Password Manager CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  add <service> <username> <password>   Add a new password")
	fmt.Println("  list                                 List all saved accounts")
	fmt.Println("  search <service>                     Search by service")
}