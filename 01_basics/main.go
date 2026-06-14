package main

import ( 
    "fmt"
    "math"
)

func main() {
    fmt.Println("Hello Go!");
    age, name:=20, "dhiraj";
    const ab=23;
    // ab=42;
    fmt.Println(age);
    fmt.Println("square of 5",math.Pow(5,2));
    fmt.Println("name:",name);

    islogged:=true;
    issubscription:=!false
    fmt.Println(islogged && issubscription);

    marks:=81;
    if(marks>90){
        fmt.Println("score is A+");
    } else if(marks>80){
        fmt.Println("score is A");
    }else if(marks>70){
        fmt.Println("score is B+");
    } else if(marks>60){
        fmt.Println("score is B");
    }else {
        fmt.Println("score is poor")
    }
    
}