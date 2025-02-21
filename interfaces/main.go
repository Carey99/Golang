// interface - a type that lists methods without providing its code
package main

import "fmt"

// Defining an interface
type Animals interface {
	Makesound()
}

// defining makesound method in the interface
func (d Dog) Makesound() {
	fmt.Println("Woof!")
}

//Implementing interface with a dog struct

type Dog struct {
	Name string
}

// Polymorphism in Golang interface
// Allows different structs to be treated as an instance of interface
// Add a cat struct
type Cat struct {
	Name string
}

func (c Cat) Makesound() {
	fmt.Println("Meow")
}

func main() {
	//usage
	var Mydog Animals = Dog{Name: "Buddy"}

	Mydog.Makesound()

	animals := []Animals{Dog{Name: "Buddy"}, Cat{Name: "Whiskers"}}

	for _, animal := range animals {
		animal.Makesound()
	}
}
