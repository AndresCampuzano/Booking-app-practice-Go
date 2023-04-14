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

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for name
	fmt.Println("Please enter your name:")
	fmt.Scan(&firstName)

	// ask user for last name
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	// ask user for email
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	// ask user for number of tickets
	fmt.Println("Please enter the number of tickets:")
	fmt.Scan(&userTickets)

	// print user information
	fmt.Printf("Hello %v %v, welcome to %v\n", firstName, lastName, conferenceName)
	fmt.Printf("You have booked %v tickets\n", userTickets)
	fmt.Printf("A confirmation email has been sent to %v\n", email)
}
