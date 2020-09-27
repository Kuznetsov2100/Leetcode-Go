/*
 * @lc app=leetcode.cn id=1413 lang=golang
 *
 * [1413] 逐步求和得到正数的最小值
 */

// @lc code=start
func minStartValue(nums []int) int {
	sum, minimum := 0, math.MaxInt64
	for _, num := range nums {
		sum += num
		if sum < minimum {
			minimum = sum
		}
	}
	if minimum > 0 {
		return 1
	}
	return 1-minimum
}
// @lc code=end

