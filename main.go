package main

import "fmt"

func main(){
	const appName = "GoBooking"
	var eventName = "Go Conference"
	const ticketsTotal uint8 = 50
	var ticketsSold uint8 = 0
	var ticketsRemaining uint8 = ticketsTotal - ticketsSold

	fmt.Println("Welcome to our booking app", appName)
	fmt.Printf("Get your ticket here for the event: %s\n", eventName)
	fmt.Printf("we have %d of %d tickets available\n",ticketsRemaining, ticketsTotal )

	var firstName string
	var lastName string
	var email string
	var ticketsBought uint8
  
	fmt.Println("Would you please tell me your first name:")
	fmt.Scan(&firstName)
	fmt.Println("And your last name:")
	fmt.Scan(&lastName)
	fmt.Println("What's your email address:")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want to buy:")
	fmt.Scan(&ticketsBought) 
	ticketsSold = ticketsBought
	ticketsRemaining = ticketsTotal - ticketsSold
   
	fmt.Printf("Thank you %s %s for booking %d tickets. Tickets will be sent to your email address: %s\n", firstName,lastName, ticketsBought, email)
	fmt.Printf("We now have %d tickets remaining for %s\n", ticketsRemaining,eventName)
}
