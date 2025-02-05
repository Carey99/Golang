package main

import "fmt"

// define the struct
type Person struct {
	Name string
	Age  int
	Arr  []int
}

func main() {

	//initialize the struct
	var p1 Person
	p1.Name = "Carey"
	p1.Age = 20
	p1.Arr = []int{1, 2, 3, 4}

	fmt.Println(p1)

	//another way to initialiize
	p2 := Person{
		Name: "Edwin",
		Age:  10,
		Arr:  []int{34, 35},
	}

	fmt.Println(p2)

	//using struct literals to initialize
	p3 := Person{"Akida", 34, []int{200, 40, 60}}
	fmt.Println(p3)
}
