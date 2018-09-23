package main

import (
	"fmt"
	"strconv"
)

// Person struct (aka class)
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)

func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age) + " years old."
}

// hasBirthday method (pointer receiver)

func (person *Person) hasBirthday() {
	person.age++
}

// getMarried method (pointer receiver)

func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	personOne := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative

	personTwo := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(personOne)

	// fmt.Println(personOne.firstName)

	personOne.hasBirthday()
	personOne.getMarried("Williams")

	personTwo.getMarried("Thompson")

	fmt.Println(personTwo.greet())
}
