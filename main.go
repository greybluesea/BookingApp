package main

import (
	"fmt"
	"strings"
)

// package level variables
const appName = "GoBooking"
const eventName = "Go Conference"
const ticketsTotal uint8 = 50

var ticketsSold uint8 = 0
var ticketsRemaining uint8
var bookings []string

func main() {

	greet()

	for {
		// variables for each customer
		var firstName string
		lastName := ""
		email := ""
		var ticketsBought uint8

		// interaction for user input
		for {
			fmt.Print("Would you please tell me your first name:")
			fmt.Scan(&firstName)

			if len(firstName) >= 2 {
				break
			} else {
				fmt.Println("Please enter a valid name that is no less than 2 characters")
				continue
			}
		}

		for {
			fmt.Print("And your last name:")
			fmt.Scan(&lastName)

			if len(lastName) >= 2 {
				break
			} else {
				fmt.Println("Please enter a valid name that is no less than 2 characters")
				continue
			}
		}

		for {
			fmt.Print("What's your email address:")
			fmt.Scan(&email)

			if strings.Contains(email, "@") {
				break
			} else {
				fmt.Println("Please re-enter a valid email address that contains '@'")
				continue
			}
		}

		for {
			fmt.Print("How many tickets do you want to buy:")
			fmt.Scan(&ticketsBought)
			if ticketsBought <= 0 {
				fmt.Println("The number of tickets you want to buy has to be a positive number")
				continue
			} else if ticketsBought > ticketsRemaining {
				fmt.Printf("we have only %d tickets left, please re-enter\n", ticketsRemaining)
				continue
			} else {
				break
			}
		}

		// confirmation of user input
		fmt.Printf("Thank you %s %s for booking %d tickets. Tickets will be sent to your email address: %s\n", firstName, lastName, ticketsBought, email)
		bookings = append(bookings, firstName)

		// count tickets
		countTickets(ticketsBought)

		// count bookings
		countBookings(firstName)

		// check for ending
		if ticketsRemaining == 0 {
			fmt.Printf("Tickets for %s have been all sold out. Come back next year! \n", eventName)
			break
		}
	}
}

func greet() {
	fmt.Printf("\tWelcome to our booking app: %s!\n", appName)
	fmt.Printf("Get your ticket here for the event: %s\n", eventName)
	fmt.Printf("we have %d tickets in total available\n", ticketsTotal)
}

func useFirstNameBookings(bookings []string) []string {
	var firstNameBookings = []string{}
	for _, booking := range bookings {
		var fullName = strings.Fields(booking)
		firstNameBookings = append(firstNameBookings, fullName[0])
	}
	return firstNameBookings
}

func countTickets(ticketsBought uint8) {
	ticketsSold = ticketsSold + ticketsBought
	ticketsRemaining = ticketsTotal - ticketsSold
	fmt.Printf("We now have %d tickets remaining for %s.\n", ticketsRemaining, eventName)
}

func countBookings(firstName string) {
	bookings = append(bookings, firstName)
	firstNameBookings := useFirstNameBookings(bookings)
	fmt.Printf("Our current booking customers are: %v \n", firstNameBookings)
}
