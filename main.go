package main

import (
	"fmt"
	"strings"
)

// Package level variables
// Allows variables to be available to all functions
// Requires var syntax
var conferenceName = "Go Conference"

const conferenceTickets int = 50

// uint cannot be negative
var remainingTickets uint = 50

// Slice is created without a fixed size
var bookings = []string{}

// declare the entry point
func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		// Collect the return values from the function
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first name of each attendee: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked. Come back next year.")
				// Terminates the for loop
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Enter a valid email address.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}
}

// Parameters require a type
func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Return type must be specified
func getFirstNames() []string {
	firstNames := []string{}
	// For slices and arrays, ranges provides the index and value for each element
	// For loop use an _ as a blank identifier or variable not used
	for _, booking := range bookings {
		// Fields splits the string with white space
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// Go allows functions to return multiple values
// Must specify return type of each value
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// Each value that is being returned
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for their name
	// Set a pointer
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)
}
