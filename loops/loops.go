package main

import "fmt"

func main() {
	//x := 0

	//for x < 5 {
	//	fmt.Println(x)
	//	x++
	//}

	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}

	arr := []int{20, 30, 40, 50}
	//for index, value := range arr {
	//	fmt.Printf("Index at %v is %v \n", index, value)
	//}

	for _, value := range arr {
		fmt.Printf("Value is %v \n", value)
	}
}
