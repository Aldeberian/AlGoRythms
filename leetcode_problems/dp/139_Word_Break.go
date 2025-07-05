package leetcode

import "strings"

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

// Note that the same word in the dictionary may be reused multiple times in the segmentation.

func wordBreak(s string, wordDict []string) bool {
	save := make(map[string]int)
	var rec func(stri string, word string) bool

	rec = func(stri string, word string) bool {
		if stri == word {
			save[word] = 2
			return true
		}

		if strings.HasPrefix(stri, word) {
			for _, val := range wordDict {
				if save[stri[len(word):]] == 1 {
					continue
				} else if save[stri[len(word):]] == 2 {
					save[stri[len(word):]] = 2
					return true
				}

				if rec(stri[len(word):], val) {
					save[stri[len(word):]] = 2
					return true
				}
			}
		}

		save[word] = 1
		return false
	}

	return rec(s, "")
}
