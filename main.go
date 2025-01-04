package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
)

//custom functions

func add(x int, y int) int{
	return x+y;
}

func swap(x,y string)(string,string){
	return y,x;
}

func split(sum int)(x,y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

//Variables
var c, python, java bool; // default value is false

//data types

var(
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Println("Hello World");
	fmt.Println("My favourite number is : ",rand.Intn(10)); // prints a random number between 0 and 9

	fmt.Println("Addition of 2 and 3  is : ",add(2,3))
	fmt.Println(swap("Hello","World"))
	fmt.Println(split(17))
	fmt.Println(c,python,java)

	name := "Deepanshu"
	fmt.Println(name)

	fmt.Printf("Type %T Value : %v\n",ToBe,ToBe)
	fmt.Printf("Type %T Value : %v\n",MaxInt,MaxInt)
	fmt.Printf("Type %T Value : %v\n",z,z)

	//type conversion
	var x,y int = 3,4;
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("Type %T Value : %v\n",f,f);
	var m uint = uint(f);
	fmt.Printf("Type %T Value : %v\n",m,m);

	//loops
	sum:=0;

	for i:=0;i<10;i++{
		sum+=i;
		fmt.Println(sum);
	}
	fmt.Println(sum);

	//switch

	switch os:= runtime.GOOS;os{
	case "darwin":
		fmt.Println("OS X");
	case "linux":
		fmt.Println("Linux");
	default:
		//freebsd,openbsd,windows
		fmt.Printf("%s.\n",os);
	}
}
