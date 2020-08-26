/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1 // 上下左右四个边界点
	for up < down && left < right {
		for i := left; i < right; i++ { // 上边
			res = append(res, matrix[up][i])
		}
		for i := up; i < down; i++ { // 右边
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- { // 下边
			res = append(res, matrix[down][i])
		}
		for i := down; i > up; i-- { // 左边
			res = append(res, matrix[i][left])
		}
		// 四个边界点收缩一圈
		up++
		down--
		left++
		right--
	}
	if up == down { // 剩下最后一行
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
	} else if left == right { // 剩下最后一列
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][left])
		}
	}
	return res
}
// @lc code=end

