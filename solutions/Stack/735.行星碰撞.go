/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		boom := false
		for len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid < 0 {
			a, b := abs(stack[len(stack)-1]), abs(asteroid)
			if a == b {
				stack = stack[:len(stack)-1]
				boom = true
				break
			} else if a < b {
				stack = stack[:len(stack)-1]
			} else if a > b {
				boom = true
				break
			}		
		}	
		if !boom {
			stack = append(stack, asteroid)
		}	
	}
	return stack
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

