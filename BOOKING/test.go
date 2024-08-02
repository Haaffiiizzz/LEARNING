package main

import (
	"fmt"
)

func main() {
	var firstName string
	fmt.Println("Enter your first name: ")
	_, err := fmt.Scan(&firstName)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Hello, ", firstName)
}
