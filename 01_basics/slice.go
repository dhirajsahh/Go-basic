package main

import (
	"fmt"
	"math"
)

func main() {

	// result:=[]int{1,2,3}
	// result[0]=1;
	// result[2]=2;
	// result=append(result,20,20)
	// fmt.Println(result)

	arr1 := []int{10, 20, 30}
	arr2 := []int{40, 50}
	arr3 := append(arr2, arr1...)
	fmt.Println(arr1, arr2, arr3)

	result := 95000 * math.Pow(1.1, 6)
	fmt.Println(result)
}
