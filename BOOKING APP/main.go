package main

import "fmt"

func main() {
	var name = "Go Conference"
	const tickets = 50
	var ticketsLeft = 50

	fmt.Println("Welcome to the", name, "booking application")
	fmt.Println("There are a total of", tickets, "tickets. There are", ticketsLeft, "left.")

	var userName string
	var userTickets int

	fmt.Scan(&userName)

	fmt.Println(&ticketsLeft)

	userName = "Fix"
	userTickets = 2
	fmt.Println(userName, userTickets)
}
