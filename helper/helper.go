package helper

import (
	"strings"
)

// Turn validateUserInput into helper function
// Function is exported when first letter is capitalized
// ValidateUserInput

// Go allows functions to return multiple values
// Must specify return type of each value
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// Each value that is being returned
	return isValidName, isValidEmail, isValidTicketNumber
}
