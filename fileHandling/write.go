package main

import (
	"os"
)

func main() {

	file, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// b := []byte("hello world")
	// n, err := file.Write(b)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n)
	file.WriteString("Hello world1")
	file.WriteString("Hello world2")

}
