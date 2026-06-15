/*
package main

import "fmt"

func main(){

	// person:="dhiraj"
	// person="ram"
	// person:="hari";
	// fmt.Println(person)

	person:=map[string]string{
		"name":"Dhiraj",
		"address":"ktm",
	}
	person["clz"]="wrc"
	person["add"]="Siraha"
	fmt.Println(person["name"],person["address"],person["clz"],person["add"])
}
*/

package main
import "fmt"

func main(){

	person:=map[string]string{
		"name":"Dhiraj",
		"addr":"Siraha",
		"clz":"wrc",
	}
	person["level"]="Bachelor"
	vn,okn:=person["name"]
	vs,oks:=person["abc"];
	fmt.Printf("%T",vs)
	fmt.Println(vn,okn);
	fmt.Println(vs,oks);
    if oks {
		fmt.Println("Hello oks presenet",oks)
	}else {
		fmt.Println("oks not present",vs)
	}
	fmt.Println(person)

	points:=map[string]int{
		"a":10,
		"b":5,
	}
	va,oka:=points["a"];
	vc,okc:=points["c"]
	fmt.Println(va,oka)
	fmt.Println(vc,okc)

	points1:=map[string]float32{
		"a":10,
		"b":5,
	}
	va1,oka1:=points1["a"];
	vc1,okc1:=points1["c"]

	fmt.Println(va1,oka1)
	fmt.Println(vc1,okc1)

	details:=map[string]interface{}{
		"name":"Dhiraj",
		"addr":"Siraha",
		"class":10,
		"roll":11,
		"islogged":false,
		"marks": []int{20,30,90},
		
	}
	for key,value:=range details{
		fmt.Println(key,value);
	}
}