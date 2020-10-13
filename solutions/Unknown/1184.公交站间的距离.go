/*
 * @lc app=leetcode.cn id=1184 lang=golang
 *
 * [1184] 公交站间的距离
 */

// @lc code=start
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var totalSum, sum int
	begin := min(start, destination)
	end := max(start, destination)
	for i, dis := range distance {
		totalSum += dis
		if i >= begin && i < end {
			sum += dis
		}
	}
	return min(sum, totalSum-sum)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

