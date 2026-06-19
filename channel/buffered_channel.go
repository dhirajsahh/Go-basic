package main

import (
	"fmt"
)

// func emailSender(emailChan chan string, done chan bool)
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChan {
		fmt.Println("sending email to ", email)
	}
}
func main() {

	emailChan := make(chan string, 30)
	done := make(chan bool)

	go emailSender(emailChan, done)
	for i := 0; i < 30; i++ {

		emailChan <- fmt.Sprintf("%v@gmail.com", i)
	}
	// to show that email chan is non blocking here because it is a buffered channael
	fmt.Println("Hello world")
	close(emailChan)
	<-done
}
