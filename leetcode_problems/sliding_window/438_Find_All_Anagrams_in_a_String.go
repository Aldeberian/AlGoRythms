package leetcode

import (
	"unicode"
)

// Given two strings s and p, return an array of all the start indices of p's in s. You may return the answer in any order.

func stringToInt(s string) [26]int {
	var res [26]int

	for _, val := range s {
		res[int(unicode.ToUpper(val)-'A')]++
	}

	return res
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	arrayP := stringToInt(p)
	startArray := stringToInt(s[0:len(p)])
	res := []int{}

	if arrayP == startArray {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		startArray[int(unicode.ToUpper(rune(s[i-len(p)]))-'A')]--
		startArray[int(unicode.ToUpper(rune(s[i]))-'A')]++

		if arrayP == startArray {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}
