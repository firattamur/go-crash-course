package main

import "fmt"

func main() {

	x := 5
	y := 10

	// if - else
	if x < y {
		fmt.Printf("%d < %d\n", x, y)
	} else {
		fmt.Printf("%d > %d\n", x, y)
	}

	// if - else if - else
	color := "blue"

	if color == "green" {
		fmt.Println("Green!")
	} else if color == "red" {
		fmt.Println("Red!")
	} else {
		fmt.Println("Blue!")
	}

	// switch
	switch color {

	case "red":
		fmt.Println("Red!")

	case "blue":
		fmt.Println("Blue!")

	case "green":
		fmt.Println("Green!")

	}

}
