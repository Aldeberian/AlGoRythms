package leetcode

// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

//     '?' Matches any single character.
//     '*' Matches any sequence of characters (including the empty sequence).

// The matching should cover the entire input string (not partial).

func isMatch(s string, p string) bool {
	sIndex, pIndex := 0, 0

	for {
		if (sIndex == len(s) && pIndex != len(p)) || (sIndex != len(s) && pIndex == len(p)) {
			return false
		}

		if s[sIndex] == p[pIndex] || p[pIndex] == '?' {
			sIndex++
			pIndex++
			if sIndex == len(s) && pIndex == len(p) {
				return true
			}
		}

		if p[pIndex] == '*' {
			if pIndex == len(p)-1 {
				return true
			}
			for {
				if p[pIndex+1] == s[sIndex] {
					pIndex++
					break
				}
				sIndex++
				if sIndex == len(s) {
					return false
				}
			}
		}
	}
}
