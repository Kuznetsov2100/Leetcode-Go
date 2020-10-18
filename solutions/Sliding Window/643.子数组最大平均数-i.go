/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	left, right, windowsum := 0, 0, 0
	max := -100001.0
	for right < len(nums) {
		windowsum += nums[right]
		right++
		if right - left + 1 > k {
			if ans := float64(windowsum)/float64(k); ans > max {
				max = ans
			}
			windowsum -= nums[left]
			left++	
		}
	}
	return max
}
// @lc code=end

