package leetcode

// For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

func gcdOfStrings(str1 string, str2 string) string {
	word := []rune{}

	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] == str2[i] {
			if len(str1) > len(str2) && len(word)+1 < len(str1)*2 {
				word = append(word, rune(str1[i]))
			} else if len(str2) > len(str1) && len(word)+1 < len(str2)*2 {
				word = append(word, rune(str1[i]))
			}
		}
	}

	return string(word)
}
