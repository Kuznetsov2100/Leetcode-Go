/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else if v == 0 {
			count = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}
// @lc code=end

