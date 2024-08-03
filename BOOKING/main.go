package main

import (
	"fmt"
	"strings"
)

// comments in go
var conferenceName = "Go Conference"

const tickets int = 50

var ticketsLeft uint = 50
var bookings []string

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getInput()
		isValidName, isValidEmail, isValidTicket := validateInput(firstName, lastName, email, userTickets)
		if isValidName && isValidEmail && isValidTicket {
			bookTickets(userTickets, firstName, lastName, email)

			var firstNames = getFirstNames()
			fmt.Printf("First names of bookings are %v:\n", firstNames)

			if ticketsLeft == 0 {
				fmt.Println("Run out of tickets. Please come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is invalid.")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid.")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets is invalid.")
			}
		}
	}
}

func greetUser() {
	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("There are a total of", tickets, "tickets. There are", ticketsLeft, "left.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	ticketsLeft = ticketsLeft - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You'll get a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v.\n", ticketsLeft, conferenceName)
}
