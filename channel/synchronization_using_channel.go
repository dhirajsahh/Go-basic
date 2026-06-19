// package main

// import "fmt"

// func task(done chan bool) {

// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Processing....")
// }

// func main() {

// 	done := make(chan bool)

// 	go task(done)

// 	<-done

// }

package main

import (
	"fmt"
)

func Task(i int, done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println(i)
}

func main() {

	done := make(chan bool)
	a := 2

	// for i := 0; i <= 10; i++ {
	// 	go Task(i)
	// }
	go Task(a, done)
	<-done
	//time.Sleep(time.Second * 3)

}
