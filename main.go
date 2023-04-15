package main

import (
	"fmt"
)

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets remaining and %v tickets are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend.")

	// Bookings structure
	type Bookings struct {
		FirstName string
		LastName  string
		Email     string
		Tickets   int
	}

	var bookings []Bookings

	for remainingTickets > 0 {

		var firstName string
		var lastName string
		var email string
		var userTickets int

		// Gets the user information
		fmt.Println("Please enter your name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email:")
		fmt.Scan(&email)

		fmt.Println("Please enter the number of tickets:")
		fmt.Scan(&userTickets)

		// Updates remaining tickets
		remainingTickets = remainingTickets - userTickets

		// Adds a new user object to bookings slice
		bookings = append(bookings, Bookings{firstName, lastName, email, userTickets})

		// Prints user information
		fmt.Printf("Hello %v %v, welcome to %v\n", firstName, lastName, conferenceName)
		fmt.Printf("You have booked %v tickets\n", userTickets)
		fmt.Printf("A confirmation email has been sent to %v\n", email)

		// Prints remaining tickets
		fmt.Printf("We have %v tickets remaining\n", remainingTickets)

		// Prints total bookings
		fmt.Printf("We have %v bookings\n", len(bookings))
	}

	// Prints all bookings
	fmt.Println("All bookings:", bookings)
}
