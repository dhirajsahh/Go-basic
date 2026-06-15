
package main

import "fmt"

func greet(name string) int{
	fmt.Println("name",name);
	return 5;
}
func main(){
  
	a:=greet("Dhiraj")
	fmt.Println(a)
	
}