package main

import "fmt"

func main(){
	names := [4]string{"John", "Paul", "George", "Ringo"};

	fmt.Println(names);

	a := names[0:2];
	b := names[1:3];

	fmt.Println(a, b);

	b[0] = "XXX";

	fmt.Println(a, b);
	fmt.Println(names);

	var s []int;
	printSlice(s);

	s = append(s, 0);
	printSlice(s);

	s = append(s, 1);
	printSlice(s);

	s = append(s, 2, 3, 4);

	printSlice(s);

	for i,v := range s{
		fmt.Printf("%d %d\n", i, v);
	}
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s);
}
