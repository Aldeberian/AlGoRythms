package leetcode

//Given an integer numRows, return the first numRows of Pascal's triangle.

func Generate(numRows int) [][]int {
	response := [][]int{}

	for i := 0; i < numRows; i++ {
		temp := []int{}

		temp = append(temp, 1)
		if i == 0 {
			response = append(response, temp)
			continue
		}

		for j := 1; j < i; j++ {
			temp = append(temp, response[i-1][j-1] + response[i-1][j])
		}

		temp = append(temp, 1)
		response = append(response, temp)
	}

	return response
}