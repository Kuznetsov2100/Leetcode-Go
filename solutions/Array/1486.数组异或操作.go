/*
 * @lc app=leetcode.cn id=1486 lang=golang
 *
 * [1486] 数组异或操作
 */

// @lc code=start
func xorOperation(n int, start int) int {
	res := start
	for i := 0; i < n; i++ {
		numsi := start+2*i
		if i > 0 {
			res ^= numsi
		}
	}
	return res
}
// @lc code=end

