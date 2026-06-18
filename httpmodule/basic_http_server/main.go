package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get is allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write([]byte("Hello from Go net/http server"))
}

func main() {
	//defining route
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("go to port 8000")
	err := http.ListenAndServe(":8000", nil)
	fmt.Println(err)
}
