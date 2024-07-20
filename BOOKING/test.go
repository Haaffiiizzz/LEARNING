package main

import "fmt"

func main() {
	var input string

	fmt.Print("Please enter some text: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("You entered:", input)
}
