/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	left, right := 0, n
	var res int
	for left <= right {
		mid := left + (right - left)/2
		if (1+mid)*mid/2 <= n {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}
	return res
}
// @lc code=end

