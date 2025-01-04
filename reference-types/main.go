package main;

import "fmt"

func main(){
	i,j := 42,2701;

	p := &i;
	fmt.Println(*p); // point to i
	*p = 21; // set i through the pointer->dereferencing
	fmt.Println(i); // see the new value of i

	p = &j; // point to j
	*p = *p / 37; // divide j through the pointer
	fmt.Println(j); // see the new value of j
}
