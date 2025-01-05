package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex;

func main() {
	m = make(map[string]Vertex)
	m["A"] = Vertex{1, 2}
	m["B"] = Vertex{3, 4}
	m["C"] = Vertex{5, 6}

	fmt.Println(m["A"]);
	fmt.Println(m);

	n  := make(map[string]int);
	n["Answer"] = 42;

	fmt.Println(n);
	fmt.Println("The value:", n["Answer"]);

	delete(n, "Answer");
	fmt.Println(n);

	v,ok := n["Answer"];
	if ok{
		fmt.Println("The value:",v);
	}else{
		fmt.Println("Not found");
		fmt.Println(n);
	}

	//loop over the map

	for key,value := range n{
		fmt.Printf("%s : %d\n",key,value);
	}
}
