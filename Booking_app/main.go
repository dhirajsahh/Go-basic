package main

import (
	"booking_app/greatuser"
	"booking_app/userinput"
	"fmt"
	"strings"
	"time"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50

	var remainingTickets uint = 50
	var bookings = []string{}

	//fmt.Printf("conference Name is %T and conference Tickets is  %T \n", conferenceName, conferenceTickets)

	// fmt.Println("Welcome to", conferenceName, "booking application")

	// fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, "are still available.")

	greatuser.GreatUser(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := userinput.GetUserInput()

		isValidUserName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets)
		if !isValidUserName {
			fmt.Println("Too short userName")
			break
		}
		if !isValidEmail {
			fmt.Println("invalid email")
			break
		}
		if !isValidUserTickets {
			fmt.Println("Tickets should be grater 0 and less than 50")
		}
		bookings = append(bookings, firstName+" "+lastName)

		firstNames := []string{}
		for _, value := range bookings {
			fname := strings.Fields(value)
			firstNames = append(firstNames, fname[0])
		}
		fmt.Println(firstNames)

		remainingTickets = remainingTickets - uint(userTickets)
		go sendTicket(firstName, lastName, email, userTickets)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	}

}

func sendTicket(firstName, lastName, email string, userTickets int) {

	time.Sleep(10 * time.Second)
	fmt.Println("Before sending ticket ####")
	ticketDetails := fmt.Sprintf("Thank You %v %v for booking ticket with us we have sent you the ticket details at %v and you have booked %v tickets ", firstName, lastName, email, userTickets)
	fmt.Println(ticketDetails)
	fmt.Println("After sending Ticket ####")

}
