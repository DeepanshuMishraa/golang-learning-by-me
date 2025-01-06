package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	var num int

	fmt.Printf("Enter the number: \n")
	fmt.Scanf("%d", &num)

	solution := fact(num)

	fmt.Printf("The factorial of %v is %v \n", num, solution)
}
