package main;

import "fmt"


// create a struct

type Vertex struct{
	x int
	y int
}

func main(){
	//initialise a struct

	v:= Vertex{1,2}
	v.x = 4;
	fmt.Println(v.x)
	fmt.Println(v);

	p := &v;
	p.x = 1e9;
	fmt.Println(v);

}
