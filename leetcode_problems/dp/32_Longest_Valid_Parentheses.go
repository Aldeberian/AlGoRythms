package leetcode

import (
	utils "AlGoRythms/utils"
	"fmt"
	"slices"
)

// Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses

// Il faudrait sauvegarder le dernier endroit où on a eu une paranthèse ouvrante et calculer le max depuis ici si left != 0

func LongestValidParentheses(s string) int {
	stack := utils.Stack[int]{}
	save := []int{}

	for i, val := range s {
		if val == '(' {
			stack.Push(i)
		} else {
			if !stack.IsEmpty() {
				left, _ := stack.Pop()

				save = append(save, left)
				save = append(save, i)
			}
		}
	}
	slices.Sort(save)
	fmt.Println(save)
	max, tempMax := 0, 0

	for i, val := range save {
		if i == 0 || val == save[i-1]+1 {
			tempMax++
		} else {
			if tempMax > max {
				max = tempMax
			}
			tempMax = 1
		}
	}
	if tempMax > max {
		max = tempMax
	}
	return max
}
