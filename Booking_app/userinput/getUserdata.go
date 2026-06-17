package userinput

import "fmt"

func GetUserInput() (string, string, string, int) {
	var firstName, lastName, email string
	var userTickets int
	fmt.Println("Enter your first Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
