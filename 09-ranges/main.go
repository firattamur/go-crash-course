package main

import "fmt"

func main() {

	ids := []int{33, 2, 3, 43, 34}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	emailv2 := map[string]string{"Bob": "Bob@example.com", "John": "John@example.com"}

	for k, v := range emailv2 {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emailv2 {
		fmt.Printf("Name: %s\n", emailv2[k])
	}

}
