package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	dirInfo, err := f.ReadDir(-1)

	for _, di := range dirInfo {
		fmt.Println(di.Name())
	}

}
