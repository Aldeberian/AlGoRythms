package leetcode

// Given a string s, return true if the s can be palindrome after deleting at most one character from it.

func ValidPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			var gauche bool
			if left > 0 {
				gauche = isValid(s[:left-1] + s[left:])
			} else {
				gauche = isValid(s[left:])
			}
			return gauche || isValid(s[:right]+s[right+1:])
		}
		left++
		right--
	}

	return true
}

func isValid(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
