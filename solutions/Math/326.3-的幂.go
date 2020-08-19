/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n == 3 {
		return true
	}
	if n % 3 != 0 {
		return false
	}
	return isPowerOfThree(n/3)
}
// @lc code=end

