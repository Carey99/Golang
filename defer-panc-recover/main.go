package main

import "fmt"

// 'defer' keyword is used to delay execution
// 'panic' stops normal execution immediately
// recover is used to catch a panic and prevent the program from crashing
func check(age int) {
	if age < 0 {
		panic("value below zero encountered : age cannot be under 0")
	}
	fmt.Println(age)
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	fmt.Println("Starting the program")
	panic("Something went wrong!")
	fmt.Println("This won't be printed")
}
