package main

import (
	"fmt"
)

func main() {
	// Define a map

	// emails := make(map[string]string)

	// Assing key : values

	// emails["bob"] = "bob@gmail.com"
	// emails["sharon"] = "sharon@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	// Declare map and add key : values

	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com", "mike": "mike@gmail.com"}

	fmt.Println(len(emails))

	// Delete a key

	delete(emails, "bob")

	fmt.Println(len(emails))
}
