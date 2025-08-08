package leetcode

// We are given an array asteroids of integers representing asteroids in a row. The indices of the asteriod in the array represent their relative position in space.

// For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

// Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	stack := [][2]int{}
	removed := make(map[int]bool)

	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack = append(stack, [2]int{asteroids[i], i})
		} else {
			if len(stack) > 0 {
				if stack[len(stack)-1][0]+asteroids[i] > 0 {
					removed[i] = true
				} else if stack[len(stack)-1][0]+asteroids[i] == 0 {
					removed[stack[len(stack)-1][1]] = true
					removed[i] = true
					stack = stack[:len(stack)-1]
				} else {
					removed[stack[len(stack)-1][1]] = true
					i--
				}
			}
		}
	}

	for i := range asteroids {
		if !removed[i] {
			res = append(res, asteroids[i])
		}
	}

	return res
}