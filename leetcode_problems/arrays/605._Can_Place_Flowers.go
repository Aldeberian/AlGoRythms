package leetcode

// You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

func canPlaceFlowers(flowerbed []int, n int) bool {
	res := 0

	for i := 0; i < len(flowerbed); i++ {
		if i == 0 && (len(flowerbed) == 1 || flowerbed[1] == 0) && flowerbed[0] == 0 {
			res++
			flowerbed[i] = 1
		} else if i != 0 && i == len(flowerbed)-1 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
			flowerbed[i] = 1
			res++
		} else if i != 0 && flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			res++
		}
	}

	return res >= n
}
