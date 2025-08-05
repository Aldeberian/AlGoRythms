package leetcode

import "math"

// TODO

// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

// You have the following three operations permitted on a word:

//     Insert a character
//     Delete a character
//     Replace a character

// Greedy version

// func minDistance(word1 string, word2 string) int {
// 	storage := make(map[string]int)

// 	var rec func(word string) int

// 	rec = func(word string) int {
// 		if word == word2 {
// 			return 0
// 		}

// 		if _, ok := storage[word]; ok {
// 			return storage[word]
// 		}

// 		mini := math.MaxFloat64

// 		if len(word) > len(word2) {
// 			for i := range word {
// 				if i >= len(word2) || word[i] != word2[i] {
// 					if i < len(word)-1 {
// 						temp := float64(1 + rec(word[:i]+word[i+1:]))
// 						mini = min(mini, temp)
// 					} else {
// 						temp := float64(1 + rec(word[:i]))
// 						mini = min(mini, temp)
// 					}
// 				}
// 			}
// 		} else if len(word) == len(word2) {
// 			for i := range word {
// 				if word[i] != word2[i] {
// 					tempword := []rune(word)
// 					tempword[i] = rune(word2[i])
// 					newWord := string(tempword)
// 					temp := float64(1 + rec(newWord))
// 					mini = min(mini, temp)
// 				}
// 			}
// 		} else {
// 			for i := range word2 {
// 				if i > len(word)-1 {
// 					tempword := word + string(word2[i])
// 					temp := float64(1 + rec(tempword))
// 					mini = min(mini, temp)
// 				} else if word[i] != word2[i] {
// 					tempword := word[:i] + string(word2[i]) + word[i+1:]
// 					temp := float64(1 + rec(tempword))
// 					mini = min(mini, temp)
// 				}
// 			}
// 		}

// 		storage[word] = int(mini)

// 		return int(mini)
// 	}

// 	return rec(word1)
// }

// Not greedy but not accurate

func minDistance(word1 string, word2 string) int {
	storage := make(map[string]int)

	var rec func(word string) int

	rec = func(word string) int {
		if word == word2 {
			return 0
		}

		if _, ok := storage[word]; ok {
			return storage[word]
		}

		mini := math.MaxFloat64

		if len(word) > len(word2) {
			for i := range word {
				if i >= len(word2) || word[i] != word2[i] {
					if i < len(word)-1 {
						return 1 + rec(word[:i]+word[i+1:])
					} else {
						return 1 + rec(word[:i])
					}
				}
			}
		} else if len(word) == len(word2) {
			for i := range word {
				if word[i] != word2[i] {
					tempword := []rune(word)
					tempword[i] = rune(word2[i])
					newWord := string(tempword)
					return 1 + rec(newWord)
				}
			}
		} else {
			for i := range word2 {
				if i > len(word)-1 {
					tempword := word + string(word2[i])
					temp := float64(1 + rec(tempword))
					mini = min(mini, temp)
				} else if word[i] != word2[i] {
					tempword := word[:i] + string(word2[i]) + word[i+1:]
					return 1 + rec(tempword)
				}
			}
		}

		storage[word] = int(mini)

		return int(mini)
	}

	return rec(word1)
}
