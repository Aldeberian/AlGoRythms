package leetcode

// Given a string s, reverse only all the vowels in the string and return it.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

func reverseVowels(s string) string {
	stack := []rune{}

	for _, l := range s {
		switch l {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			stack = append(stack, l)
		}
	}

	res := []rune{}
	for _, l := range s {
		switch l {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		default:
			res = append(res, l)
		}
	}

	return string(res)
}
