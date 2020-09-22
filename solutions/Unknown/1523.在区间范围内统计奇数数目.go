/*
 * @lc app=leetcode.cn id=1523 lang=golang
 *
 * [1523] 在区间范围内统计奇数数目
 */

// @lc code=start
func countOdds(low int, high int) int {
	remainderLow, remainderHigh := low % 2, high % 2
	if remainderLow == 0 && remainderHigh == 0 {
		return helper(low, high)
	}
	if remainderLow == 1 && remainderHigh == 0 {
		return 1 + helper(low+1, high)
	}
	if remainderLow == 0 && remainderHigh == 1 {
		return 1 + helper(low, high-1)
	}
	if low == high {
		return 1
	}
	return 2 + helper(low+1, high-1)		
}

func helper(low int, high int) int {
	if low >= high {
		return 0
	}
	return (high-low)/2
}
// @lc code=end

