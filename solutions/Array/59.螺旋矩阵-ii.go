/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	cnt := 1
	up, down, left, right := 0, n-1, 0, n-1 // 上下左右四个边界点
	for up < down && left < right {
		for i := left; i < right; i++ { // 上边
			matrix[up][i]= cnt
			cnt++
		}
		for i := up; i < down; i++ { // 右边
			matrix[i][right] = cnt
			cnt++
		}
		for i := right; i > left; i-- { // 下边
			matrix[down][i] = cnt
			cnt++
		}
		for i := down; i > up; i-- { // 左边
			matrix[i][left] = cnt
			cnt++
		}
		// 四个边界点收缩一圈
		up++
		down--
		left++
		right--
	}
	if n % 2 == 1 { // 当n为奇数时,剩下最后一个点
		matrix[up][left] = cnt
	}
	return matrix
}
// @lc code=end

