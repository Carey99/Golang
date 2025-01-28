package main

import (
	"fmt"
	"strings"
)

func InitialLetters(s string) (string, string) {
	upper_s := strings.ToUpper(s)
	splitted_s := strings.Split(upper_s, " ") //[CAREY,  EDWIN]

	var res []string

	for _, v := range splitted_s {
		res = append(res, v[:1])
	}

	if len(res) > 1 {
		return res[0], res[1]
	}
	return res[0], " "
}

func main() {
	sn1, sn2 := InitialLetters("carey edwin") //Outputs: C E

	fmt.Println(sn1, sn2)
}
