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

	var userName string
	var userTickets int
	// ask user for name

	userName = "John Doe"
	userTickets = 2
	fmt.Printf("Hello %v, welcome to %v\n", userName, conferenceName)
	fmt.Printf("You have booked %v tickets\n", userTickets)

}
