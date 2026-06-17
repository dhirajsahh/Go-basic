package main

import "strings"

func validateUserInput(firstName, lastName, email string, userTickets int) (bool, bool, bool) {

	isValidUserName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= 50

	return isValidUserName, isValidEmail, isValidUserTickets
}
