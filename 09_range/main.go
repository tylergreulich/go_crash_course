package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	ids := []int{1, 2, 3, 4, 5}

	// Loop through IDs

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all IDs together

	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum:", sum)

	// Range w/ Map

	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com", "mike": "mike@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
