package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing taks", i)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
	// time.Sleep(time.Second * 2)
}
