package main

import "fmt"

func main() {

	// main go variable types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var key to define variable
	var name = "firat"
	var age int = 37

	// constant variables
	const isConst = true

	// shorthand variable declaration
	shorthand := "make it easy."
	size := 1.3

	fmt.Println(shorthand, size, isConst)
	fmt.Println(name, age)
	fmt.Printf("Type of %s: %T\n", name, name)
	fmt.Printf("Type of %d: %T\n", age, age)

}
