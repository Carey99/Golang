package main

import (
	"fmt"
	"math"
)

func Sayhello(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func Saybye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func PassingFunc(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func CircleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	//Sayhello("Carey")
	//Sayhello("Edwin")
	//Saybye("Carey")

	names := []string{"Carey", "Edwin", "Akida"}
	PassingFunc(names, Sayhello)
	PassingFunc(names, Saybye)

	a1 := CircleArea(15.5)
	a2 := CircleArea(5)

	fmt.Printf("Area a1 is %0.3f\n", a1)
	fmt.Printf("Area a2 is %0.3f\n", a2)
}
