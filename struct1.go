
package main

import "fmt"

type Person struct {

	name string
	age int
	address string
	weight float64
	isMarried bool
	hobby  []string
}
func (p Person) displayValue () {
	 fmt.Println(p.name)
	 fmt.Println(p.weight)
	 fmt.Println(p.isMarried)
	 fmt.Println(p.address)
	 fmt.Println(p.age)
	 for _,value:=range p.hobby{
		fmt.Println(value);
	}
}

func (p *Person) changeAttribute(newname string , age int){

	p.name=newname
	p.age=age
}

func main(){

	p:=Person{
		name:"Dhiraj",
		address:"kathmandu",
		age:20,
		isMarried:false,
		hobby:[]string{"playing cricket","watching movies"},
	}
     
	p.displayValue();
	p.changeAttribute("Niraj",22)
	fmt.Println("After")
	p.displayValue();
	//fmt.Println(p)


}