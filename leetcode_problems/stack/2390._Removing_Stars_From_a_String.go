package leetcode

// You are given a string s, which contains stars *.

// In one operation, you can:

//     Choose a star in s.
//     Remove the closest non-star character to its left, as well as remove the star itself.

// Return the string after all stars have been removed.

// Note:

//     The input will be generated such that the operation is always possible.
//     It can be shown that the resulting string will always be unique.

func removeStars(s string) string {
	var res []rune;
	count := 0

	tempS := []rune(s)

	for i := len(tempS)-1; i >= 0; i-- {
		if tempS[i] == '*' {
			count++
		} else if count > 0 {
			count--
		} else {
			res = append(res, tempS[i]) 
		}
	}

	for i, j := 0, len(res)-1; i < j; i,j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
