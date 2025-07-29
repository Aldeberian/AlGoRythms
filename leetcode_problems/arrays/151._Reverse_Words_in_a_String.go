package leetcode

import "strings"

// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words.
// The returned string should only have a single space separating the words. Do not include any extra spaces.

func reverseWords(s string) string {
	var stockage []string
	splitS := strings.Split(s, " ")

	for _, el := range splitS {
		temp := strings.TrimSpace(el)
		stockage = append(stockage, temp)
	}

	res := ""
	for i := len(stockage) - 1; i >= 0; i-- {
		if stockage[i] != "" {
			res += stockage[i] + " "
		}
	}

	return strings.TrimSpace(res)
}
