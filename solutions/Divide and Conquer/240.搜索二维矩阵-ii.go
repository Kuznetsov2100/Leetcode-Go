/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	// 从矩阵左下角开始遍历
	row, col := m-1, 0
	for row >= 0 && col < n {
		if val := matrix[row][col]; val > target {
			row--
		} else if val < target {
			col++
		} else if val == target {
			return true
		}
	}
	return false
}
// @lc code=end

