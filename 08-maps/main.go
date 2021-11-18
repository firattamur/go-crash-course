package main

import "fmt"

func main() {

	// define a map
	emails := make(map[string]string)

	// assign key-values
	emails["alice"] = "alice@example.com"
	emails["bob"] = "bob@example.com"
	emails["main"] = "main@example.com"

	fmt.Println(emails)

}
