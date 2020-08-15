/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	left, right, sum := 0, 0, 0
	minlen := 1 << 63 -1
	window := 0 // 记录子数组的长度
	for right < len(nums) {
		sum += nums[right]
		window++
		right++
		for sum >= s  {
			if window < minlen {
				minlen = window
			}
			sum -= nums[left]
			window--
			left++
		}
	}
	if minlen == (1 << 63 -1) {
		return 0
	}
	return minlen
}
// @lc code=end

