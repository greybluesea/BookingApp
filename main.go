package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"

	/* 	"strconv" */
	"strings"
)

// package level variables
const appName = "GoBooking"
const eventName = "Go Conference"
const ticketsTotal uint8 = 50

var ticketsSold uint8 = 0
var ticketsRemaining uint8 = ticketsTotal - ticketsSold

// var bookings = make([]map[string]string, 0)
var bookings = make([]Customer, 0)

type Customer struct {
	firstName     string
	lastName      string
	email         string
	ticketsBought uint8
}

var waitGroup = sync.WaitGroup{}

func main() {
	helper.Hello()
	greet()

	// loop through for each customer
	for {

		// take user input
		firstName, lastName, email, ticketsBought := takeUserInput()

		// confirmation of user input
		fmt.Printf("Thank you %s %s for booking %d tickets. Tickets will be sent to your email address: %s in 5 seconds.\n", firstName, lastName, ticketsBought, email)

		//construct userProfile map
		/* 	var userProfile = make(map[string]string)
		userProfile["firstName"] = firstName
		userProfile["lastName"] = lastName
		userProfile["email"] = email
		userProfile["ticketsBought"] = strconv.FormatUint(uint64(ticketsBought), 10) */

		// construct userProfile struct

		var customer = Customer{
			firstName:     firstName,
			lastName:      lastName,
			email:         email,
			ticketsBought: ticketsBought,
		}

		// send tickets
		waitGroup.Add(1)
		go sendTickets(customer)

		// count tickets
		countTickets(ticketsBought)

		// count bookings
		countBookings(customer)

		// check for ending
		if ticketsRemaining == 0 {
			fmt.Printf("Tickets for %s have been all sold out. Come back next year! \n", eventName)
			waitGroup.Wait()
			break
		}

		//break for next customer
		fmt.Println("")
		fmt.Println("")
	}
}

func greet() {
	fmt.Printf("\tWelcome to our booking app: %s!\n", appName)
	fmt.Printf("Get your ticket here for the event: %s\n", eventName)
	fmt.Printf("we have %d tickets in total available\n", ticketsTotal)
}

func takeUserInput() (string, string, string, uint8) {
	var firstName string
	lastName := ""
	email := ""
	var ticketsBought uint8

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
	return firstName, lastName, email, ticketsBought
}

func countTickets(ticketsBought uint8) {
	ticketsSold = ticketsSold + ticketsBought
	ticketsRemaining = ticketsTotal - ticketsSold
	fmt.Printf("We now have %d tickets remaining for %s.\n", ticketsRemaining, eventName)
}

func countBookings(customer /* map[string]string */ Customer) {
	bookings = append(bookings, customer)
	firstNameBookings := useFirstNameBookings(bookings)
	fmt.Printf("Our current booking customers are: %v \n", firstNameBookings)
}

func useFirstNameBookings(bookings [] /* map[string]string */ Customer) []string {
	var firstNameBookings = []string{}
	for _, booking := range bookings {
		var firstName = /*  booking["firstName"] */ booking.firstName
		firstNameBookings = append(firstNameBookings, firstName)
	}
	return firstNameBookings
}

func sendTickets(customer Customer) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("--------\n\t %s e-Tickets\n\t %d tickets for %s %s\n--------\n", eventName, customer.ticketsBought, customer.firstName, customer.lastName)
	fmt.Printf("\n\n#########\tSending Tickets \t#########\n \n%v\n#########  to email address: %s   ############\n\n", ticket, customer.email)
	waitGroup.Done()
}
