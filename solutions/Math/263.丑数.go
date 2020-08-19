/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	if num % 2 == 0 {
		return isUgly(num/2)
	} else if num % 3 == 0 {
		return isUgly(num/3)
	} else if num % 5 == 0 {
		return isUgly(num/5)
	} else {
		return false
	}
}
// @lc code=end

