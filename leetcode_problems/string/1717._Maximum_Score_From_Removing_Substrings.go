package leetcode

// You are given a string s and two integers x and y. You can perform two types of operations any number of times.

//     Remove substring "ab" and gain x points.
//         For example, when removing "ab" from "cabxbae" it becomes "cxbae".
//     Remove substring "ba" and gain y points.
//         For example, when removing "ba" from "cabxbae" it becomes "cabxe".

// Return the maximum points you can gain after applying the above operations on s.

func gain(s string, a, b byte, val int) (string, int) {
	stack := []byte{}
	score := 0
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == a && s[i] == b {
			stack = stack[:len(stack)-1]
			score += val
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack), score
}

func maximumGain(s string, x int, y int) int {
	total := 0
	if x > y {
		s, v := gain(s, 'a', 'b', x)
		total += v
		_, v = gain(s, 'b', 'a', y)
		total += v
	} else {
		s, v := gain(s, 'b', 'a', y)
		total += v
		_, v = gain(s, 'a', 'b', x)
		total += v
	}
	return total
}
