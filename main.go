package main

import (
	"fmt"
	"math/rand"
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

func main(){
	fmt.Println("Hello World");
	fmt.Println("My favourite number is : ",rand.Intn(10)); // prints a random number between 0 and 9

	fmt.Println("Addition of 2 and 3  is : ",add(2,3))
	fmt.Println(swap("Hello","World"))
}
