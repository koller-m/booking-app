package main

import (
	"fmt"
	"strconv"

	"github.com/koller-m/booking-app/helper"
)

// Package level variables
// Allows variables to be available to all functions
// Requires var syntax
var conferenceName = "Go Conference"

const conferenceTickets int = 50

// uint cannot be negative
var remainingTickets uint = 50

// Slice is created without a fixed size
// List of maps will grow dynamically
var bookings = make([]map[string]string, 0)

// declare the entry point
func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		// Collect the return values from the function
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

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
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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

	// Create an empty map
	var userData = make(map[string]string)
	// Key names can be called anything
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email

	// Convert uint to string
	// 10 is for base 10
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)
}
