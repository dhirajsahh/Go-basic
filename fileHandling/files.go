package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("file name:", fileInfo.IsDir())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file mode:", fileInfo.Mode())
	fmt.Println("file modified:", fileInfo.ModTime())

}
