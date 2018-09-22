package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// Long method

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// Fizzbuzz

	number := 100

	for j := 0; j < number; j++ {
		if j%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if j%5 == 0 {
			fmt.Println("Fizz")
		} else if j%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}
}
