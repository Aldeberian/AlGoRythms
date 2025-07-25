package leetcode

// Alice and Bob are playing a game. Initially, Alice has a string word = "a".

// You are given a positive integer k. You are also given an integer array operations, where operations[i] represents the type of the ith operation.

// Now Bob will ask Alice to perform all operations in sequence:

//     If operations[i] == 0, append a copy of word to itself.
//     If operations[i] == 1, generate a new string by changing each character in word to its next character in the English alphabet, and append it to the original word. For example, performing the operation on "c" generates "cd" and performing the operation on "zb" generates "zbac".

// Return the value of the kth character in word after performing all the operations.

// Note that the character 'z' can be changed to 'a' in the second type of operation.

// Naive approach
// func kthCharacter(k int64, operations []int) byte {
// 	suite := []byte{'a'}

// 	for _, op := range operations {
// 		if op == 0 {
// 			suite = append(suite, suite...)
// 		} else {
// 			n := len(suite)
// 			for i := 0; i < n; i++ {
// 				suite = append(suite, byte(((int(suite[i]-'a')+1)%26)+'a'))
// 			}
// 		}
// 	}

// 	return suite[k]
// }

func kthCharacter(k int64, operations []int) byte {
	shift := 0
	n := len(operations)
	lastIndex := 0

	val := 1
	for i := range n {
		val *= 2
		if val >= int(k) {
			lastIndex = i
			break
		}
	}

	for i := lastIndex; i >= 0; i-- {
		half := int64(1) << i

		if k > half {
			k -= half
			if operations[i] == 1 {
				shift++
			}
		}
	}

	return byte((shift % 26) + 'a')
}
