/*
 * @lc app=leetcode.cn id=1266 lang=golang
 *
 * [1266] 访问所有点的最小时间
 */

// @lc code=start
func minTimeToVisitAllPoints(points [][]int) int {
	// 遍历整个数组，对于数组中的相邻两个点，计算出它们的切比雪夫距离，所有的距离之和即为答案。
	var res int
	for i := 0; i < len(points)-1; i++ {
		res += max(abs(points[i][0]-points[i+1][0]), abs(points[i][1]-points[i+1][1]))
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

