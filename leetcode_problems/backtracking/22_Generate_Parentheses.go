package leetcode

import (
	"strings"
)

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

func GenerateParenthesis(n int) []string {
	var response []string

	var rec func(rang int, chaine string)

	rec = func(rang int, chaine string) {
		if rang == n*2 {
			response = append(response, chaine)
			return
		}

		left := strings.Count(chaine, "(")
		right := strings.Count(chaine, ")")

		if left > right {
			rec(rang+1, chaine+string(')'))
		}
		if left-right+rang < n*2 {
			rec(rang+1, chaine+string('('))
		}
	}

	rec(0, "")

	return response
}
