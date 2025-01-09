package main

import "fmt"

func main() {
	numbers := []int{1, 2,2,4,5,6,7,8,9,10,5}

	seen:= make(map[int]bool);
	hasDuplicates := false;

	for _,num:=range numbers{
		if seen[num]{
			hasDuplicates = true;
			break;
		}

		seen[num] = true;
	}

	if hasDuplicates{
		fmt.Println("Has duplicates")
	}else{
		fmt.Println("No duplicates")
	}
	
}
