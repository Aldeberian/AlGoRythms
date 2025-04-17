package main

import (
	"fmt"
	"unicode"
)

func stringToInt(s string) [26]int {
	var res [26]int

	for _, val := range s {
		res[int(unicode.ToUpper(val)-'A')]++
	}

	return res
}

func main() {
	fmt.Println(stringToInt("test"))
}
