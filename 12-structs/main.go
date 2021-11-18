package main

import (
	"fmt"
	"strconv"
)

// define struct
type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int

	// also this possible
	// firstname, lastname, city, gender string
	// age                               int
}

// methods should be outside of the structs
func (p Person) greet() string {
	return "Hello, " + p.firstname + "Your Age: " + strconv.Itoa(int(p.age))
}

// pointer receiver function
func (p *Person) hasBirtday() {
	p.age++
}

func main() {

	person1 := Person{firstname: "firat",
		lastname: "tamur",
		city:     "istanbul",
		gender:   "male",
		age:      23}

	person2 := Person{"bob", "alice", "new york", "female", 23}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstname)

	fmt.Println(person1.greet())
	person1.hasBirtday()
	fmt.Println(person1.greet())

}
