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

	greetUser(conferenceName, tickets, ticketsLeft)

	firstName, lastName, email, userTickets := getInput()

	for {

		isValidName, isValidEmail, isValidTicket := validateInput(firstName, lastName, email, userTickets, ticketsLeft)
		if isValidName && isValidEmail && isValidTicket {

			bookTickets(ticketsLeft, userTickets, bookings, firstName, lastName, email, conferenceName)

			var firstNames = getFirstNames(bookings)
			fmt.Printf("First names of Bookings are %v: \n", firstNames)

			if ticketsLeft == 0 {
				fmt.Println("Run out of tickets. Please co me back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is invalid.")
			}
			if !isValidEmail {
				fmt.Print("Email is invalid.")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets is invalid.")
			}

		}
	}

}

func greetUser(conferenceName string, tickets int, ticketsLeft uint) {
	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("There are a total of", tickets, "tickets. There are", ticketsLeft, "left.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func validateInput(firstName string, lastName string, email string, userTickets uint, ticketsLeft uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= ticketsLeft

	return isValidName, isValidEmail, isValidTicket
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

func bookTickets(ticketsLeft uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {

	ticketsLeft = ticketsLeft - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You'll get a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v.\n", ticketsLeft, conferenceName)
}
