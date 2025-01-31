//Write a function divide
//It takes two integers and returns both the quotient and the remainder
//If the divisor is 0 return 0, 0

package main

import "fmt"

func divide(a, b int) (int, int) {
	if b == 0 {
		return 0, 0
	} else {
		return a / b, a % b
	}
}

func main() {
	q, r := divide(10, 3)
	fmt.Println(q, r)
}
