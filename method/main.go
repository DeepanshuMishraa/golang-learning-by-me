package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

func(v Vertex)Abs()float64{
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long);
}

func main(){
	v := Vertex{1, 2}
	fmt.Println(v.Abs());
}
