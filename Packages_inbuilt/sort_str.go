package main

import (
	"fmt"
	"sort"
)

func main() {
	original_str := []string{"Carey", "Edwin", "Akida"}
	str := append([]string{}, original_str...)

	sort.Strings(str)
	fmt.Println(original_str)
	fmt.Println(str)

}
