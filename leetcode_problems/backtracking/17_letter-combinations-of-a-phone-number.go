package leetcode

import "strconv"

//Given a string containing digits from 2-9 inclusive,
// return all possible letter combinations that the number could represent. Return the answer in any order

func LetterCombinations(digits string) []string {
	lettres := make([]string, 10)
	lettres[2] = "abc"
	lettres[3] = "def"
	lettres[4] = "ghi"
	lettres[5] = "jkl"
	lettres[6] = "mno"
	lettres[7] = "pqrs"
	lettres[8] = "tuv"
	lettres[9] = "wxyz"

	var response []string
	var rec func(digits string, chaine string, num int)

	rec = func(digits string, chaine string, num int) {
		if num == len(digits) {
			return
		}

		index := 0
		firstDigit, _ := strconv.Atoi(string(digits[num]))
		listeLettre := lettres[firstDigit]
		for index < len(listeLettre) {
			nouvChaine := chaine + string(listeLettre[index])
			if num == len(digits)-1 {
				response = append(response, nouvChaine)
			}
			index++
			rec(digits, nouvChaine, num+1)
		}
	}

	rec(digits, "", 0)

	return response
}
