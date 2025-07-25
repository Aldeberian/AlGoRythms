package leetcode

// A fancy string is a string where no three consecutive characters are equal.

// Given a string s, delete the minimum possible number of characters from s to make it fancy.

// Return the final string after the deletion. It can be shown that the answer will always be unique.

func makeFancyString(s string) string {
	res := []rune{}

	for index, val := range s {
		if index > 1 {
			if s[index-1] != byte(val) || s[index-2] != byte(val) {
				res = append(res, val)
			}
		} else {
			res = append(res, val)
		}
	}

	return string(res)
}
