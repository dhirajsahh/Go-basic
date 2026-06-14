
package main
import ("fmt")

func main(){
    
	sum:=10;
	for i:=1;i<5;i++{
		sum+=i;
		fmt.Println(i);
	}
	fmt.Println(sum)
}