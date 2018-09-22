package main

import "fmt"

// var name = "Tyler"

func main() {
	var age = 22
	const isCool = true
	var size float32 = 1.3

	// name := "Tyler"
	// email := "tgreul@tgreul.com"

	name, email := "Tyler", "tgreul@tgreul.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
