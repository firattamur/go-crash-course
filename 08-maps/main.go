package main

import "fmt"

func main() {

	// define a map
	emails := make(map[string]string)

	// assign key-values
	emails["alice"] = "alice@example.com"
	emails["bob"] = "bob@example.com"
	emails["main"] = "main@example.com"

	// declare and add key-value pairs
	emailv2 := map[string]string{"Bob": "Bob@example.com", "John": "John@example.com"}

	fmt.Println(emails)
	fmt.Println(emailv2)

}
