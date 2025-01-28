package main

import "fmt"

func main() {
	age := 50

	fmt.Println(age == 45)
	fmt.Println(age < 30)
	fmt.Println(age <= 35)
	fmt.Println(age > 20)
	fmt.Println(age >= 25)

	if age == 45 {
		fmt.Println("Age is 45")
	} else if age < 30 {
		fmt.Println("Age is less that 30")
	} else {
		fmt.Println("Just a normal age")
	}

	arr := []int{20, 30, 40, 50}

	for index, value := range arr {
		if index == 1 {
			fmt.Printf("Continuing from pos %v\n", index)
			continue
		}
		if index > 2 {
			fmt.Printf("Breaking at indx %v \n", index)
			break
		}
		fmt.Printf("Index %v is %v\n", index, value)
	}
}
