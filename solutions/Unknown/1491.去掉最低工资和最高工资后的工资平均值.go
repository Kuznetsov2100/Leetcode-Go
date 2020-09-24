/*
 * @lc app=leetcode.cn id=1491 lang=golang
 *
 * [1491] 去掉最低工资和最高工资后的工资平均值
 */

// @lc code=start
func average(salary []int) float64 {
	maximum, minimum := math.MinInt64, math.MaxInt64
	sum := 0
	for _, val := range salary {
		sum += val
		if val > maximum {
			maximum = val
		}
		if val < minimum {
			minimum = val
		}
	}
	return float64(sum-minimum-maximum) / float64(len(salary)-2)
}
// @lc code=end

