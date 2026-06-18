package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := map[string]any{
		"ok":   true,
		"msg":  "Hello",
		"date": time.Now().UTC(),
	}
	json.NewEncoder(w).Encode(res)
}

type User struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Age  int    `json:"age"`
}

func decodeBody(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(err)
	fmt.Println(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
func main() {

	http.HandleFunc("/", homeRoute)
	http.HandleFunc("/decode", decodeBody)
	err := http.ListenAndServe(":8002", nil)
	fmt.Println(err)
}
