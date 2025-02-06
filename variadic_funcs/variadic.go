// a function that takes multiple args of the same type
package main

import "fmt"

// uses elpsis for definition
func sum(numbers ...int) int {
	total := 0

	//numbers will treated as slice
	for _, num := range numbers {
		total += num
	}
	return total
}

//variadic func must be the last parameter in a function
func greeting(prefix string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s %s\n", prefix, name)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))

	//passing a slice into a variadic function
	fmt.Println(sum([]int{10, 30}...))

	greeting("Hello", "Carey", "Edwin", "Akida")
}
