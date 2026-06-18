package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only get method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// res, err := w.Write([]byte("welcome and go to /test?name=abc"))
	// fmt.Println("res", res)
	// fmt.Println(err)
	http.ServeFile(w, r, "http_multiple_routes/index.html")

}
func testHandler(w http.ResponseWriter, r *http.Request) {

	urlParams := r.URL.Query()
	name := (urlParams.Get("name"))
	fmt.Println(urlParams)

	_, _ = w.Write([]byte(name))
}
func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/test", testHandler)
	err := http.ListenAndServe(":8001", nil)
	fmt.Println(err)
}
