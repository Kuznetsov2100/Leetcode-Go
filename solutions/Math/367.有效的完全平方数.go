/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 2, num
	for left <= right {
		mid := left + (right - left)/2
		if pow := mid * mid; pow == num {
			return true
		} else if pow < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
// @lc code=end

