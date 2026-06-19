// package main

// import (
// 	"fmt"
// )

// func receive(messageChan chan string) {

// 	msg := <-messageChan
// 	//msg1 := <-messageChan

// 	fmt.Println(msg)
// 	//fmt.Println(msg1)
// }
// func add(messageChan chan int, a int, b int) {
// 	sum := a + b
// 	messageChan <- sum
// }
// func main() {

// 	messageChan := make(chan int)
// 	go add(messageChan, 5, 4)
// 	sum := <-messageChan

// 	fmt.Println(sum)

// 	// messageChan := make(chan string)

// 	// go receive(messageChan)messsageChan
// 	// messageChan <- "ping"
// 	// // for {
// 	// // 	messageChan <- "ping"
// 	// // }
// 	// time.Sleep(time.Second * 2)
// }

package main

import "fmt"

func receive(messageChan chan int) {
	msg := <-messageChan
	fmt.Println(msg)
}

func send(sumChan chan int, a, b int) {
	sum := a + b
	sumChan <- sum

}
func main() {
	//sumChan := make(chan int)
	var sumChan chan int
	sumChan = make(chan int)
	go send(sumChan, 4, 5)
	sum := <-sumChan
	fmt.Println(sum)

	// messageChan := make(chan int)
	// go receive(messageChan)
	// messageChan <- 10

	// msg := <-messageChan
	// fmt.Println(msg)
}
