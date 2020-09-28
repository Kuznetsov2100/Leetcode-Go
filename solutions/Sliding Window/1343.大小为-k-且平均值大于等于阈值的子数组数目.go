/*
 * @lc app=leetcode.cn id=1343 lang=golang
 *
 * [1343] 大小为 K 且平均值大于等于阈值的子数组数目
 */

// @lc code=start
func numOfSubarrays(arr []int, k int, threshold int) int {
	left, right, res, windowsum := 0, 0, 0, 0
	for right < len(arr) {
		windowsum += arr[right]
		right++
		if right-left + 1 > k {
			if float64(windowsum)/float64(k) >= float64(threshold) {
				res++
			}
			windowsum -= arr[left]
			left++
		}
	}
	return res
}
// @lc code=end

