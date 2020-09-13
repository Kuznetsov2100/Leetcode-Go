/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)
	i, j, flag := 0, 0, -1
	for {
		res = append(res, matrix[i][j])
		for flag == -1 && i > 0 && j < n-1 || flag == 1 && i < m-1 && j > 0 {
			i, j = i+flag, j-flag 
			res = append(res, matrix[i][j])
		}
		if (i == 0 || i == m-1) && j < n-1 {
			j++
		} else if (j == 0 || j == n-1) && i < m-1 {
			i++
		} else {
			break
		}
		flag = -flag
	}
	return res
}
// @lc code=end

