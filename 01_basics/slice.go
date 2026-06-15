package main

import "fmt"

func main(){

	result:=[]int{1,2,3}
	result[0]=1;
	result[2]=2;
	result=append(result,20,20)
	fmt.Println(result)
}

