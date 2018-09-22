package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not red or blue")
	}

	switch color {
	case "blue":
		fmt.Println("Color is blue")
	case "red":
		fmt.Println("Color is red")
	default:
		fmt.Println("Color is not blue or red")
	}
}
