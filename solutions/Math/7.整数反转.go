/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	origin := x
	if x < 0 {
		x = -x
	}
	reverse := 0
	for x != 0 {
		reverse = reverse*10 + x % 10
		x = x / 10
	}
	if origin >= 0 {
		if reverse > (1 << 31 - 1) {
			return 0
		}
		return reverse
	} else {
		if -reverse < (-1 << 31) {
			return 0
		}
		return -reverse
	}
}
// @lc code=end

