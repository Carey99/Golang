// Do a FizzBuzz challenge looping upto 20
// if n % 3 == 0 --> print Fizz
// if n % 5 == 0 --> print Buzz
// if n % (3 and 5) == 0 --> print FizzBuzz

package main

import "fmt"

func fizzbuzz() {
	for val := 1; val <= 20; val++ {
		if val%3 == 0 && val%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if val%3 == 0 {
			fmt.Println("Fizz")
		} else if val%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(val)
		}
	}
}
