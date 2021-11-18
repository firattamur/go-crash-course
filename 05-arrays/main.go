package main

import "fmt"

func main() {

	// arrays has fixed length
	var fruits [2]string

	// assign values
	fruits[0] = "apple"
	fruits[1] = "orange"

	// declare and assign
	fruits2 := []string{"apple", "orange"}

	// print arrays and elements
	fmt.Println(fruits, fruits2)
	fmt.Println(fruits[0])

	// slices no fixed size
	fruitsSlice := []string{"apple", "orange", "grape"}

	fmt.Println(fruitsSlice)
	fmt.Println(len(fruitsSlice))
	fmt.Println(fruitsSlice[:1])

}
