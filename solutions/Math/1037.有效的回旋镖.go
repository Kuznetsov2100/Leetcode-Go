/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
func isBoomerang(points [][]int) bool {
	if (points[0][0] == points[1][0] && points[0][1] == points[1][1] ) ||
	(points[0][0] == points[2][0] && points[0][1] == points[2][1]) {
		return false		
	}
	dx1 := points[0][0]-points[1][0]
	dy1 := points[0][1]-points[1][1]
	dx2 := points[1][0]-points[2][0]
	dy2 := points[1][1]-points[2][1]
	if dx2 * dy1 != dx1* dy2 {
		return true
	}
	return false
}
// @lc code=end

