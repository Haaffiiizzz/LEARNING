package helper

import "strings"

func ValidateInput(firstName string, lastName string, email string, userTickets uint, ticketsLeft uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= ticketsLeft

	return isValidName, isValidEmail, isValidTicket
}
