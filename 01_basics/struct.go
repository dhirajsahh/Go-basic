package main

import "fmt"


type Address struct{

	name string
	address string
	education string
	percentage int
	roll int
}
func main (){
	var a= Address{"Dhiraj","wrc","bachelor",80,11};
	a.roll=12;
	fmt.Println(a);
   
	a2:=Address{"yubraj","trichandra","bachelor",62,0}
	fmt.Println(a2)

	a3:=Address{name:"yp",address:"bishnu"}
	fmt.Println(a3)
}