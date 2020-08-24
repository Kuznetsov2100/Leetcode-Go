/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	// 找到左侧第一个 citations[i] >= n - i
	var res int
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right - left)/2
		if citations[mid] >= n - mid {
			res = n - mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
// @lc code=end

