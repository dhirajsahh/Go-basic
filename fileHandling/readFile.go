package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size()

	b := make([]byte, fileSize)
	n, err := f.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	for i := 0; i < len(b); i++ {
		fmt.Printf(string(b[i]))
	}
	fmt.Println()

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
