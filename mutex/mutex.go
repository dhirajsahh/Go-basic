package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu    sync.Mutex
	//name  string
}

func (p *Post) inc(wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
		p.mu.Unlock()

	}()
	//p.name = name
	p.mu.Lock()
	p.views++
}
func main() {
	myPost := Post{
		views: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}
	//myPost.inc()
	//myPost.inc("dhiraj")
	wg.Wait()

	fmt.Println(myPost.views)

}
