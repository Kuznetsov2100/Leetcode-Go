/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
func smallestRangeI(A []int, K int) int {
	maximum := math.MinInt64
	minimum := math.MaxInt64
	for _, num := range A {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
	}
	max := maximum - K
	min := minimum + K
	if max-min < 0 {
		return 0
	}
	return max-min
}
// @lc code=end

