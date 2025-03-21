package leetcode

// Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses

// Il faudrait sauvegarder le dernier endroit où on a eu une paranthèse ouvrante et calculer le max depuis ici si left != 0

func LongestValidParentheses(s string) int {
	left, tempMax, max := 0, 0, 0

	for _, val := range s {
		if val == '(' {
			left++
		} else {
			if left > 0 {
				left--
				tempMax += 2
				if tempMax > max {
					max = tempMax
				}
			} else {
				tempMax = 0
			}
		}
	}
	return max
}
