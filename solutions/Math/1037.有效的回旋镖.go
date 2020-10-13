/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
func isBoomerang(points [][]int) bool {
	// 根据斜率来判断， 将除法转换成叉乘，可避免除数为0以及点是否相同的判断
	dx1 := points[0][0]-points[1][0]
	dy1 := points[0][1]-points[1][1]
	dx2 := points[1][0]-points[2][0]
	dy2 := points[1][1]-points[2][1]
	if dx2 * dy1 == dx1 * dy2 {
		return false
	}
	return true
}
// @lc code=end

