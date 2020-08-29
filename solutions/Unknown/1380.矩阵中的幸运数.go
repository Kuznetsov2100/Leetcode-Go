/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */

// @lc code=start
func luckyNumbers (matrix [][]int) []int {
	var res []int
	rows := make([]int,len(matrix))
	for i := range rows {
		rows[i] = 1 << 31-1
	}
	cols := make([]int, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			val := matrix[i][j]
			rows[i] = min(rows[i], val)
			cols[j] = max(cols[j], val)
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			val := matrix[i][j]
			if rows[i] == val && cols[j] == val {
				res = append(res, val)
			}
		}
	}
	return res
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

