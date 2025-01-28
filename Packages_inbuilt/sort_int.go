package main

import (
	"fmt"
	"sort"
)

func main() {
	original := []int{20, 23, 45, 40}
	slices := append([]int{}, original...)

	sort.Ints(slices) //sort modifies the slice directly
	fmt.Println(original)
	fmt.Println(slices)
}
