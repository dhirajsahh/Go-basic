package main

import "fmt"

func main(){

	var x int =5;
	p:=&x;
	fmt.Println(p)

	fmt.Println(*p)

	*p=*p+10;
	fmt.Println(*p);
	fmt.Println(x);

}