package main

import "fmt"

// IndexInt returns the index of the first occurrence of x in s, or -1 if not present.
func IndexInt(s []int, x int) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// IndexString returns the index of the first occurrence of x in s, or -1 if not present.
func IndexString(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

//Now using Generics

// Index returns the index of the first occurrence of x in s, or -1 if not present.

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{1, 2, 3, 4, 5}
	fmt.Println(Index(si, 3))

	ss := []string{"a", "b", "c", "d", "e"}
	fmt.Println(Index(ss, "c"))

}
